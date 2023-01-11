package main

import (
	"fmt"

	"github.com/rodert/hepburn/config"
	"github.com/rodert/hepburn/internal/version"
	"github.com/rodert/hepburn/web"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	configPath = ""
)

func webCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "web",
		Short: "run web service",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("logo...")

			if err := config.NewConfigure(configPath); err != nil {
				logrus.Errorf(configPath)
				logrus.Fatal(err)
				return
			}
			web.Run()
		},
	}
}

func main() {
	root := &cobra.Command{Use: "hepburn", Version: version.VersionInfo()}
	// 指定解析的配置文件
	root.PersistentFlags().StringVarP(&configPath, "configure", "c", "./config.hcl", "configure file path")

	root.AddCommand(webCmd())
	if err := root.Execute(); err != nil {
		logrus.Error(err)
	}
}
