package command

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
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
	expire    uint8
	randomPwd bool
	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "Upload data to pastebin.",
		Args:  cobra.MaximumNArgs(1),
		RunE:  upload,
	}
	pool = []uint8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
)

func UploadCommand() *cobra.Command {
	uploadCmd.Flags().Bool("help", false, "help for upload")
	uploadCmd.Flags().StringVarP(&password, "private-share-password", "P", "", " Optional. Private share with specificated password.")
	uploadCmd.Flags().BoolVarP(&randomPwd, "private-share", "p", false, "Optional. Private share. Will using a random password for private share.")
	uploadCmd.Flags().Uint8Var(&expire, "e", 24, "Optional. Set to 0 means burn-after-read. Default 24. (unit: hrs)")
	return uploadCmd
}

func upload(command *cobra.Command, args []string) error {
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
	} else {
		return
	}
}

func uploadToPasteBin(context []byte) (err error) {
	_, err = fmt.Fprintf(os.Stderr, "todo: upload context: "+string(context))
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(request)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(response) // <- do not forget to release

	request.Header.SetMethod(fasthttp.MethodPost)
	request.Header.Set("X-Real-IP", "1.1.1.1")
	request.SetRequestURI(host + "/api/upload")
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
	_, err = field.Write([]byte(strconv.Itoa(int(expire))))
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
	fmt.Fprintf(os.Stderr, "Server Response:\n")
	fmt.Fprintf(os.Stderr, "Http Status Code: %d\n", response.StatusCode())
	fmt.Fprintf(os.Stderr, "Http Response Body:\n")
	fmt.Printf(string(response.Body()))
	fmt.Fprintf(os.Stderr, "\n")
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
	buffer := make([]byte, 4)
	for i := 0; i < 4; i++ {
		buffer[i] = pool[rand.Intn(len(pool))]
	}
	result := string(buffer)
	fmt.Fprintf(os.Stderr, "Private share password: %v\n", result)
	return result
}
