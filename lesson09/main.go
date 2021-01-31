package main

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// logFile, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer logFile.Close()
	// log.SetOutput(logFile)

	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs2.log",
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     2,
		Compress:   true,
	})
	for i := 0; i < 100000; i++ {
		log.Println("this is a test log", i)
	}
}
