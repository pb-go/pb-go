package clipkg

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/user"
	"strings"
)

// define the config file path and init for root command
var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "pb-cli",
		Short: "Client for pb-go.",
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringP("host", "H", "", "pb-go service url")
	err := viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	rootCmd.AddCommand(UploadCommand(), GetCommand(), DeleteCommand(), StatusCommand())
	rootCmd.SetHelpCommand(nil)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		current, err := user.Current()
		if err != nil {
			log.Fatal(err.Error())
		}
		viper.AddConfigPath(current.HomeDir)
		viper.SetConfigName(".pbcli")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error for reading config file: ~/.pbcli.yaml")
	}
	_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
}

// AcquireValidGlobalFlag : validation global flag
func AcquireValidGlobalFlag() {
	// global flag and config validation
	host := viper.Get("host").(string)
	if !(len(host) == 0) && !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		_, _ = fmt.Fprintln(os.Stderr, "Invalid host url:"+host)
		log.Fatalln("Host should start with http:// or https://")
	}
}

// Execute : so-called init function
func Execute() error {
	return rootCmd.Execute()
}
