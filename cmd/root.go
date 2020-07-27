package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/biosugar0/postmaster/config"
	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	appName   = "p8r"
	shortDesc = "A graphical CLI for your AWS SNS/SQS management."
	longDesc  = "p8r is a CLI to view and manage for your AWS SNS/SQS."
)

var (
	version, commit, date = "dev", "dev", "n/a"
	flags                 *config.Flags
	awsConfig             *session.Options

	rootCmd = &cobra.Command{
		Use:   appName,
		Short: shortDesc,
		Long:  longDesc,
		RunE:  run,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd())
	initFlags()

	rootCmd.Flags().SortFlags = false
	rootCmd.Flags().SetInterspersed(false)
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic().Err(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	fmt.Println("hello")
	spew.Dump(flags.Profile)

	return nil
}

func loadConfiguration() *config.Config {
	log.Info().Msg("P8r starting up...")
	// TODO: configuration
	return &config.Config{}
}
func initFlags() {
	flags = config.NewFlags()
	rootCmd.Flags().StringVarP(
		flags.Profile,
		"profile", "p",
		config.DefaultProfile,
		"Use a specific profile from your AWS credential file",
	)
}
