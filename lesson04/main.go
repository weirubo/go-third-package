package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io"
	"log"
	"os"
)

func main() {
	// 创建一个监听器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	// 关闭监听器
	defer watcher.Close()
	// 开始监听事件
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//log.Println("event:", event)
				//if event.Has(fsnotify.Create) {
				//	log.Println("Create：", event.Name)
				//}
				if event.Has(fsnotify.Write) {
					log.Println("Write：", event.Name)
					// 自动加载文件内容
					f, _ := os.Open("log.txt")
					_, _ = io.Copy(os.Stdout, f)
				}
				//if event.Has(fsnotify.Remove) {
				//	log.Println("Remove：", event.Name)
				//}
				//if event.Has(fsnotify.Rename) {
				//	log.Println("Rename：", event.Name)
				//}
				//if event.Has(fsnotify.Chmod) {
				//	log.Println("Chmod：", event.Name)
				//}
				//case err, ok := <-watcher.Errors:
				//	if !ok {
				//		return
				//	}
				//	log.Println("error:", err)
			}
		}
	}()
	// 添加监听目录
	err = watcher.Add("./")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("this is a test output")
	fmt.Println("222 this is a test output")
	// 永久阻塞 main goroutine
	<-make(chan struct{})
}
