package clipkg

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
)

var (
	password  string
	randomPwd bool
	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "Upload data to pastebin.",
		Args:  cobra.MaximumNArgs(1),
		RunE:  uploadOnline,
	}
	pool = []uint8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
)

// UploadCommand : sub-command processing function for upload
func UploadCommand() *cobra.Command {
	uploadCmd.Flags().Bool("help", false, "help for upload")
	uploadCmd.Flags().StringVarP(&password, "private-share-password", "P", "", " Optional. Private share with specificated password.")
	uploadCmd.Flags().BoolVarP(&randomPwd, "private-share", "p", false, "Optional. Private share. Will using a random password for private share.")
	uploadCmd.Flags().UintP("expire", "e", 24, "Optional. Set to 0 means burn-after-read. Default 24. (unit: hrs)")
	_ = viper.BindPFlag("expire", uploadCmd.Flags().Lookup("expire"))
	viper.SetDefault("expire", 24)
	return uploadCmd
}

// the first param will always be *cobra.Command, please do not delete it even if not used in function.
func uploadOnline(command *cobra.Command, args []string) error {
	AcquireValidGlobalFlag()
	if len(args) == 0 {
		stdin, err := readFromStdin()
		if err != nil {
			return err
		}
		return uploadToPasteBin(stdin)
	} else {
		file, err := readFromFile(args[0])
		if err != nil {
			return err
		}
		return uploadToPasteBin(file)
	}
}

func readFromFile(path string) (context []byte, err error) {
	return ioutil.ReadFile(path)
}

func readFromStdin() (context []byte, err error) {
	context, err = bufio.NewReader(os.Stdin).ReadBytes(0)
	if err == io.EOF {
		return context, nil
	}
	return
}

func uploadToPasteBin(context []byte) (err error) {
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(request)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(response) // <- do not forget to release

	request.Header.SetMethod(fasthttp.MethodPost)
	request.Header.Set("X-Real-IP", "1.1.1.1")
	request.SetRequestURI(viper.Get("host").(string) + "/api/upload")
	writer := multipart.NewWriter(request.BodyWriter())
	// setup password
	if password != "" || randomPwd {
		field, err := writer.CreateFormField("p")
		if err != nil {
			log.Fatal(err.Error())
		}
		_, err = field.Write([]byte(fetchPassword()))
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	// setup expire
	field, err := writer.CreateFormField("e")
	if err != nil {
		log.Fatal(err.Error())
	}

	expire := viper.Get("expire")
	var result string
	switch t := expire.(type) {
	case int:
		result = strconv.Itoa(t)
	case string:
		result = t
	default:
		result = fmt.Sprintf("%v", t)
	}

	_, err = field.Write([]byte(result))
	if err != nil {
		log.Fatal(err.Error())
	}

	// setup context
	field, err = writer.CreateFormField("d")
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = field.Write(context)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = writer.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = writer.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	request.Header.SetContentType(writer.FormDataContentType())
	err = fasthttp.Do(request, response)
	if err != nil {
		log.Fatalln("Connect to Server Error. Check your config please.")
	}
	fmt.Println(" If you need raw format data, just append `f=raw` as your snippet URL param.")
	_, _ = fmt.Fprintf(os.Stderr, "Server Response:\n")
	_, _ = fmt.Fprintf(os.Stderr, "Http Status Code: %d\n", response.StatusCode())
	_, _ = fmt.Fprintf(os.Stderr, "Http Response Body:\n")
	_, _ = fmt.Printf(string(response.Body()))
	_, _ = fmt.Fprintf(os.Stderr, "\n")
	return
}

func fetchPassword() string {
	if password != "" {
		return password
	} else if randomPwd {
		return generateRandomPassword()
	} else {
		log.Fatal("Should specific a password or use a random one for private share.¬")
		return ""
	}
}

func generateRandomPassword() string {
	genpwdlen := viper.GetInt("pwdlen")
	buffer := make([]byte, genpwdlen)
	for i := 0; i < genpwdlen; i++ {
		buffer[i] = pool[rand.Intn(len(pool))]
	}
	result := string(buffer)
	_, _ = fmt.Fprintf(os.Stderr, "Private share password: %v\n", result)
	_, _ = fmt.Printf("Please append `p=%v` as URL param after your snippet URL. \n", result)
	return result
}
