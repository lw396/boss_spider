package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var Config struct {
	Webhooks struct {
		Url               string
		Region            string
		ContentType       string
		Model             string
		BullBearThreshold float64
	}
}

func LoadIniConfig() {
	cfg, err := ini.ShadowLoad("config/app.cfg")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	Config.Webhooks.Url = cfg.Section("webhooks").Key("url").MustString("")
	Config.Webhooks.Region = cfg.Section("webhooks").Key("region").MustString("xiamen")
	Config.Webhooks.ContentType = cfg.Section("webhooks").Key("content_type").MustString("application/json")
	Config.Webhooks.Model = cfg.Section("webhooks").Key("model").MustString("")
	Config.Webhooks.BullBearThreshold = cfg.Section("webhooks").Key("bull_bear_threshold").MustFloat64(0.5)
}
