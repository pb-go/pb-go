package clipkg

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

type StatusData struct {
	RunHealth        int  `json:"status"`
	CaptchaEnabled   bool `json:"captcha_enabled"`
	MaxExpireTime    int  `json:"max_expire"`
	FileCheckEnabled bool `json:"abuse_detection"`
}

var (
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Fetching server running status and limitation.",
		Args:  cobra.NoArgs,
		RunE:  statusOnline,
	}
)

// GetCommand : sub-command get.
func statusCommand() *cobra.Command {
	return statusCmd
}

func statusOnline(command *cobra.Command, args []string) (err error) {

	url := viper.Get("host").(string) + "/status"
	code, body, err := fasthttp.Get(make([]byte, 0), url)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = fmt.Fprintf(os.Stderr, "Server Response:\n")
	_, _ = fmt.Fprintf(os.Stderr, "Http Status Code: %d\n", code)
	if code >= 400 {
		log.Fatalln("Your request is rejected by server. Please contact server administrator.")
	}
	_, _ = fmt.Fprintf(os.Stderr, "Http Response Body:\n")

	var serverStatusData StatusData
	err = json.Unmarshal(body, &serverStatusData)
	if err != nil {
		log.Fatalln("Please contact server administrator.")
	}
	_, _ = fmt.Printf("\t Server Status: Running")
	_, _ = fmt.Printf("\t Server Requires CAPTCHA Verification: %v \n", serverStatusData.CaptchaEnabled)
	_, _ = fmt.Printf("\t Server-allowed maximum expire time: %d (hrs) \n", serverStatusData.MaxExpireTime)
	_, _ = fmt.Printf("\t Server Enabled Abuse Detection Feature: %v \n", serverStatusData.FileCheckEnabled)
	_, _ = fmt.Fprintf(os.Stderr, "\n")

	return
}
