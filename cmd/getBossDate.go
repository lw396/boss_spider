package cmd

import (
	"boss_spider/service"
	"github.com/spf13/cobra"
)

var getBossDate = &cobra.Command{
	Use:   "getBossDate",
	Short: "获取Boss直聘职位信息",
	Run: func(cmd *cobra.Command, args []string) {
		service.GetBossDate()
		return
	},
}

func init() {
	rootCmd.AddCommand(getBossDate)
}
