package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func Run() {
	var command = &cobra.Command{
		Use:   "Fullstack Go Post HTMX App",
		Short: "i use this app to learn fullstack go",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(webCmd())

	if err := command.Execute(); err != nil {
		log.Fatal().Msgf("failed to execute command, err: %v", err.Error())
	}
}
