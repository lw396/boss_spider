package cmd

import "github.com/spf13/cobra"

var getBossDate = &cobra.Command{
	Use:   "getBossDate",
	Short: "获取Boss直聘职位信息",
}

func init() {
	rootCmd.AddCommand(getBossDate)
}
