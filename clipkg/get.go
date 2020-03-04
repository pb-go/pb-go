package clipkg

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

var (
	snipPassword string
	getCmd       = &cobra.Command{
		Use:   "get",
		Short: "Fetching data from patesbin with id.",
		Args:  cobra.MinimumNArgs(1),
		RunE:  getOnline,
	}
)

// GetCommand : sub-command get.
func GetCommand() *cobra.Command {
	getCmd.Flags().StringVarP(&snipPassword, "password", "p", "", "Optional. Provide password for private share.")
	return getCmd
}

func getOnline(command *cobra.Command, args []string) (err error) {

	request := MakeRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	url := viper.Get("host").(string) + "/" + args[0]

	url += "?f=raw"
	if snipPassword != "" {
		url += "&p=" + snipPassword
	}

	request.Header.SetMethod(fasthttp.MethodGet)
	request.SetRequestURI(url)
	err = fasthttp.Do(request, response)
	code := response.StatusCode()
	body := response.Body()

	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = fmt.Fprintf(os.Stderr, "Server Response:\n")
	_, _ = fmt.Fprintf(os.Stderr, "Http Status Code: %d\n", code)
	if code >= 400 {
		_, _ = fmt.Fprintf(os.Stderr, "Your request is rejected by server. Please check your password or Snippet ID.")
	}
	_, _ = fmt.Fprintf(os.Stderr, "Http Response Body:\n")
	_, _ = fmt.Printf(string(body))
	_, _ = fmt.Fprintf(os.Stderr, "\n")

	return
}
