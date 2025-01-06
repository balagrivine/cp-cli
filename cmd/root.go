package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cp-cli",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return copyFile(args)
	},
}
	return cmd
}

func Execute() {
	cmd := rootCmd()
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
