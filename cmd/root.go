package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spider",
	Short: "启动爬虫程序",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("请输入相应子命令")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
