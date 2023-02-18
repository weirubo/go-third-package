package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"time"
)

var sugaredLogger *zap.SugaredLogger
var loggerV2 *zap.Logger

func main() {
	InitLogger()
	defer sugaredLogger.Sync()
	//InitLoggerV2()
	//defer loggerV2.Sync()
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run()
}

func ping(c *gin.Context) {
	sugaredLogger.Debug("call func ping")
	//loggerV2.Debug("call func ping", zap.Int("code", 200))
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func InitLogger() {
	//core := zapcore.NewCore(enc(), ws(), enab())
	//core := zapcore.NewCore(enc(), wsV2(), enab())
	core := zapcore.NewCore(enc(), wsV3(), enab())
	logger := zap.New(core)
	sugaredLogger = logger.Sugar()
}

func InitLoggerV2() {
	core := zapcore.NewCore(enc(), ws(), enab())
	loggerV2 = zap.New(core)
}

func enc() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.TimeKey = "time"
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	return zapcore.NewJSONEncoder(cfg)
}

func ws() zapcore.WriteSyncer {
	logFileName := fmt.Sprintf("./%v.log", time.Now().Format("2006-01-02"))
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Fatal(err)
	}
	return zapcore.AddSync(logFile)
}

func wsV2() zapcore.WriteSyncer {
	return zapcore.NewMultiWriteSyncer(ws(), zapcore.AddSync(os.Stdout))
}

func wsV3() zapcore.WriteSyncer {
	logFileName := fmt.Sprintf("./%v.log", time.Now().Format("2006-01-02"))
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   false,
	}
	return zapcore.AddSync(lumberjackLogger)
}

func enab() zapcore.LevelEnabler {
	return zapcore.DebugLevel
}
