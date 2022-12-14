package main

import (
	"file_rotatelogs/core"
	"file_rotatelogs/global"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// 读取配置文件，创建全局配置viper对象
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	// 创建日志记录系统
	global.GVA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)

	global.GVA_LOG.Info("nihao")
}
