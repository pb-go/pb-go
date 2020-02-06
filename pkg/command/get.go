package command

import "github.com/spf13/cobra"

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "Get",
		Short: "Fetching data from patesbin with id.",
	}
}
