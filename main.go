/*
 * @Author: NyanCatda
 * @Date: 2022-07-16 15:11:05
 * @LastEditTime: 2022-07-16 15:27:59
 * @LastEditors: NyanCatda
 * @Description: 主文件
 * @FilePath: \SharePointProxy\main.go
 */
package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"strconv"

	"github.com/McShare/SharePointProxy/Internal/Config"
	"github.com/McShare/SharePointProxy/Internal/Constant"
	"github.com/McShare/SharePointProxy/Web/Routers"
	"github.com/gin-gonic/gin"
	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/AyaLog/ModLog/GinLog"
)

func main() {
	// 参数解析
	RunPort := flag.Int("port", 8000, "指定运行端口")
	ConfigPath := flag.String("config", "./config.yml", "指定配置文件路径")
	ColorPrint := flag.Bool("ColorPrint", true, "输出彩色文本")
	flag.Parse()

	// 设置日志参数
	AyaLog.ColorPrint = *ColorPrint
	AyaLog.LogPath = Constant.LogPath

	// 设置配置文件路径
	Config.ConfigPath = *ConfigPath

	// 加载配置文件
	if err := Config.LoadConfig(); err != nil {
		AyaLog.Error("System", err)
		AyaLog.Error("System", errors.New("请检查配置文件"))
		return
	}

	if Config.Get.Run.Mode == "release" {
		AyaLog.LogLevel = AyaLog.INFO
	}

	AyaLog.Info("System", Constant.Codename+"启动，当前版本为: "+Constant.Version+"，运行模式: "+Config.Get.Run.Mode)

	// 发行模式下调整Gin模式
	if Config.Get.Run.Mode == "release" {
		// Gin调整模式
		gin.SetMode(gin.ReleaseMode)
		// 关闭默认的日志输出
		gin.DefaultWriter = ioutil.Discard
	}

	// 初始化GIN
	r := gin.Default()

	// 注册日志组件
	r.Use((GinLog.GinLog()))

	// 拆分注册路由
	Routers.SetupRouter(r)

	// 运行
	AyaLog.Info("System", Constant.Codename+"启动完成！正在监听端口："+strconv.Itoa(*RunPort))
	if err := r.Run(":" + strconv.Itoa(*RunPort)); err != nil {
		AyaLog.Error("System", err)
		return
	}
}
