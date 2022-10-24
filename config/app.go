package config

import "goblog/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		"name": config.Env("APP_NAME", "GoBlog"),
		"env": config.Env("APP_ENV", "production"),
		"debug": config.Env("APP_DEBUG", false),
		"port": config.Env("APP_PORT", "3000"),
		"key":config.Env("APP_KEY", "7db47119a07ee57f22f67c9b1a16b9255e015d24424131e360650d40ce326120"),
	})
}