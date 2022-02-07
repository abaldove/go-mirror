package cmd

import (
	"os"
	"runtime/debug"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "crawler",
	Short: "Crawler - Backstage process for crawling URLs",
}

func Execute() {
	defer func() {
		err := recover()

		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"stack": string(debug.Stack()),
			}).Error("Captured error while executing a command!")
		}
	}()

	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
