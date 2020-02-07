package clipkg

import (
	"fmt"
	"github.com/pb-go/pb-go/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

var (
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a paste from pastebin with id.",
		RunE:  deleteOnline,
		Args:  cobra.MinimumNArgs(1),
	}
)

// DeleteCommand : Parse the delete sub-command
func DeleteCommand() *cobra.Command {
	deleteCmd.Flags().StringP("masterKey", "k", "", "Required. Master key in pb-go server's config.")
	_ = viper.BindPFlag("masterKey", deleteCmd.Flags().Lookup("masterKey"))
	return deleteCmd
}

func deleteOnline(command *cobra.Command, args []string) (err error) {
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI(viper.Get("host").(string) + "/api/admin?id=" + args[0])
	request.Header.SetMethod(fasthttp.MethodDelete)
	request.Header.Set("X-Master-Key", utils.GetUTCTimeHash(viper.Get("masterKey").(string)))

	err = fasthttp.Do(request, response)
	if err != nil {
		log.Fatal(err.Error())
	}
	respStatusCode := response.StatusCode()
	_, _ = fmt.Fprintf(os.Stderr, "Server Response:\n")
	_, _ = fmt.Fprintf(os.Stderr, "Http Status Code: %d\n", respStatusCode)
	if respStatusCode >= 400 {
		_, _ = fmt.Fprintf(os.Stderr, "Your request is rejected by server. Please check your masterkey or Snippet ID.")
	}
	_, _ = fmt.Fprintf(os.Stderr, "Http Response Body:\n")
	_, _ = fmt.Printf(string(response.Body()))
	_, _ = fmt.Fprintf(os.Stderr, "\n")
	return
}
