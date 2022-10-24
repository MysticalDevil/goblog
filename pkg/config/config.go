package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"goblog/pkg/logger"
)

var Viper *viper.Viper

// StrMap 简写 -- map[string]any
type StrMap map[string]any

func init() {
	Viper = viper.New()
	Viper.SetConfigName(".env")
	Viper.SetConfigType("env")
	Viper.AddConfigPath(".")

	err := Viper.ReadInConfig()
	logger.LogError(err)

	Viper.SetEnvPrefix("appenv")
	Viper.AutomaticEnv()
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...any) any {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return Get(envName)
}

// Add 新增配置项
func Add(name string, configuration map[string]any) {
	Viper.Set(name, configuration)
}

// Get 获取配置项，允许使用点式获取，如：app.name
func Get(path string, defaultValue ...any) any {
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...any) string {
	return cast.ToString(Get(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...any) int {
	return cast.ToInt(Get(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...any) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...any) bool {
	return cast.ToBool(Get(path, defaultValue...))
}