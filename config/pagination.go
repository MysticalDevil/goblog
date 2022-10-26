package config

import "goblog/pkg/config"

func init() {
	config.Add("pagination", config.StrMap{
		"perPage": 10,
		// URL 中以分别页数的参数
		"url_query": "page",
	})
}
