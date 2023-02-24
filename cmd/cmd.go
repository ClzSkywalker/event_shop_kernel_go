package cmd

import (
	"fmt"
	"os"

	"github.com/clz.skywalker/event.shop/kernal/cmd/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kernel",
	Short: "内核启动命令",
}

func init() {
	rootCmd.AddCommand(api.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
