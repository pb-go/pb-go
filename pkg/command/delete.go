package command

import "github.com/spf13/cobra"

func DeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a paste from pastebin with id.",
	}
}
