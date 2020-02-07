package clipkg

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/user"
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
	//rootCmd.Flags().Bool("help", false, "help for pb-cli")
	err := viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	rootCmd.AddCommand(UploadCommand(), GetCommand(), DeleteCommand())
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
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// Execute : so-called init function
func Execute() error {
	return rootCmd.Execute()
}
