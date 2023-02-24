package api

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/server"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/processx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/recoverx"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "api",
		Short:   "api startup command",
		Example: "kernel api --port=6905 --mode=release --dbPath=./assert/ectype/todo_shop.db --logPath=./logs",
		PreRun: func(cmd *cobra.Command, args []string) {
			initConfig(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			Api()
			return nil
		},
	}
)

func init() {
	Cmd.Flags().Int(constx.CmdPort, 6905, "port")
	Cmd.Flags().String(constx.CmdMode, "release", "mode")
	Cmd.Flags().String(constx.CmdDbPath, "./assert/todo_shop.db", "sqlite database path")
	Cmd.Flags().String(constx.CmdLogPath, "./logs", "log path")
}

func initConfig(cmd *cobra.Command) {
	port, _ := cmd.Flags().GetInt(constx.CmdPort)
	mode, _ := cmd.Flags().GetString(constx.CmdMode)
	dbPath, _ := cmd.Flags().GetString(constx.CmdDbPath)
	logPath, _ := cmd.Flags().GetString(constx.CmdLogPath)
	err := container.InitConfig(container.AppConfig{
		Port:    port,
		Mode:    mode,
		DbPath:  dbPath,
		LogPath: logPath,
	})
	if err != nil {
		panic(err)
	}
}

func Api() {
	go recoverx.RecoverFunc(server.CmdServer)
	processx.HandleSignal()
}
