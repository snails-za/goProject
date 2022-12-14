/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-14 21:51:45
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-14 21:52:56
 * @FilePath: /src/zap_log/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"net/http"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	simpleHttpGet("www.5lmh.com")
	simpleHttpGet("http://www.google.com")
}

func InitLogger() {
	// 指定日志写入路径
	writeSyncer := getLogWriter()
	// 指定日志编码器【如何写入内存】
	encoder := getEncoder()
	// 创建日志记录器核心【日志编码器、日志写入路径、LogLevel】
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// 生成日志记录器对象。zap.AddCaller() 用于展示调用函数信息记录到日志中【可选】
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	// 1. 创建编码器配置对象
	encoderConfig := zap.NewProductionEncoderConfig()
	// 2. 设置日志时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 3. 设置日志等级为大写，如：ERROR
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 4. NewConsoleEncoder：日志编码器
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	// 1. 如果想要追加写入可以查看我的博客文件操作那一章
	// file, _ := os.Create("./test.log")
	// return zapcore.AddSync(file)
	// 2. 在zap中加入Lumberjack支持:日志切割归档
	// Lumberjack Logger采用以下属性作为输入:
	//     Filename: 日志文件的位置
	//     MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
	//     MaxBackups：保留旧文件的最大个数
	//     MaxAges：保留旧文件的最大天数
	//     Compress：是否压缩/归档旧文件
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
