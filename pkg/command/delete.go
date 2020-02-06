package command

import (
	"fmt"
	"github.com/pb-go/pb-go/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"os"
)

var (
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a paste from pastebin with id.",
		RunE:  delete,
		Args:  cobra.MinimumNArgs(1),
	}
)

func DeleteCommand() *cobra.Command {
	deleteCmd.Flags().StringP("masterKey", "k", "", "Required. Master key in pb-go server's config.")
	viper.BindPFlag("masterKey",deleteCmd.Flags().Lookup("masterKey"))
	return deleteCmd
}

func delete(command *cobra.Command, args []string) (err error) {
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI(viper.Get("host").(string) + "/api/admin?id=" + args[0])
	request.Header.SetMethod(fasthttp.MethodDelete)
	request.Header.Set("X-Master-Key", utils.GetUTCTimeHash(viper.Get("masterKey").(string)))

	fasthttp.Do(request, response)
	err = fasthttp.Do(request, response)
	fmt.Fprintf(os.Stderr, "Server Response:\n")
	fmt.Fprintf(os.Stderr, "Http Status Code: %d\n", response.StatusCode())
	fmt.Fprintf(os.Stderr, "Http Response Body:\n")
	fmt.Printf(string(response.Body()))
	fmt.Fprintf(os.Stderr, "\n")
	return
}
