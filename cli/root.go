package cli

import (
	"github.com/lazygophers/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Run() {
	var err error

	err = rootCmd.Execute()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
}
