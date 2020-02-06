package command

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	host    string
	rootCmd = &cobra.Command{
		Use:   "pb-cli",
		Short: "Client for pb-go.",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&host, "host", "h", "", "pb-go service url")
	rootCmd.Flags().Bool("help", false, "help for pb-cli")
	rootCmd.AddCommand(UploadCommand(), GetCommand(), DeleteCommand())
}

func Execute() error {
	return rootCmd.Execute()
}
