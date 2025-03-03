package cmd

import (
	"post-htmx/internal/web"

	"github.com/spf13/cobra"
)

func webCmd() *cobra.Command {
	var port int
	var command = &cobra.Command{
		Use:   "web",
		Short: "Start frontend server",
		Run: func(cmd *cobra.Command, args []string) {
			server := web.NewServer()
			server.Run(port)
		},
	}

	command.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")
	return command
}
