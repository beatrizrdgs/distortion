package cmd

import (
	"log"
	"os"
	"os/signal"

	"github.com/beatrizrdgs/distortion/internal/distortion"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts web server",
	RunE: func(cmd *cobra.Command, args []string) error {
		svc := distortion.NewService()
		svr := distortion.NewServer("8080", svc)
		svr.Run()

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		<-stop
		log.Println("Shutting down server...")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
