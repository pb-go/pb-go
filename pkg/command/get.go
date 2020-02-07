package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"os"
)

var (
	passwordForGet string
	getCmd         = &cobra.Command{
		Use:   "get",
		Short: "Fetching data from patesbin with id.",
		Args:  cobra.MinimumNArgs(1),
		RunE:  get,
	}
)

func GetCommand() *cobra.Command {
	getCmd.Flags().StringVarP(&passwordForGet, "password", "p", "", "Optional. Provide password for private share.")
	return getCmd
}

func get(command *cobra.Command, args []string) (err error) {

	url := viper.Get("host").(string) + "/" + args[0]

	if passwordForGet != "" {
		url += "?p=" + passwordForGet
	}

	code, body, err := fasthttp.Get(make([]byte, 0), url)

	fmt.Fprintf(os.Stderr, "Server Response:\n")
	fmt.Fprintf(os.Stderr, "Http Status Code: %d\n", code)
	fmt.Fprintf(os.Stderr, "Http Response Body:\n")
	fmt.Printf(string(body))
	fmt.Fprintf(os.Stderr, "\n")

	return
}
