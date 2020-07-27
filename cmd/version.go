package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	var short bool

	command := cobra.Command{
		Use:   "version",
		Short: "Print version/build info",
		Long:  "Print version/build information",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion(short)
		},
	}

	command.PersistentFlags().BoolVarP(&short, "short", "s", false, "Prints K9s version info in short format")

	return &command
}

func printVersion(short bool) {
	const fmat = "%-20s %s\n"

	fmt.Printf(fmat, "Version", version)
	fmt.Printf(fmat, "Commit", commit)
	fmt.Printf(fmat, "Date", date)
}
