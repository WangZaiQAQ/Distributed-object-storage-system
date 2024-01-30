package main

import (
	"./heartbeat"
	"./locate"
	"./objects"
	"./temp"
	"./versions"
	"log"
	"net/http"
	"os"
)

/**
 * @Description:	起点，处理各个请求
 */

// 为什么不开4个协程执行啊？
func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)   //监听以objects开头的URL，交给后面objects包的Handler函数处理
	http.HandleFunc("/temp/", temp.Handler)         //？？？这个是干什么的啊？
	http.HandleFunc("/locate/", locate.Handler)     //
	http.HandleFunc("/versions/", versions.Handler) //获取该URL请求中的对象的所有版本的元数据
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
