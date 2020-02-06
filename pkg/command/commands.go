package command

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "pb-cli",
		Short: "Client for pb-go.",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.AddCommand(UploadCommand())
	rootCmd.AddCommand(GetCommand())
	rootCmd.AddCommand(DeleteCommand())
}

func Execute() error {
	return rootCmd.Execute()
}
