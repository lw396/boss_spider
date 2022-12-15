package service

import (
	"boss_spider/config"
	"fmt"
)

func GetBossDate() {
	config.LoadIniConfig()

	fmt.Print(config.Config.Webhooks)
}

func initItem() {
	url := config.Config.Webhooks.Url + "/" + config.Config.Webhooks.Url
	fmt.Printf("开始获取")
	fmt.Print(url)
}
