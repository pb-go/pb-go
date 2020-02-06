package command

import "github.com/spf13/cobra"

func UploadCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "Upload",
		Short: "Upload data to pastebin.",
	}
}
