package main

import (
	"embed"
	"log"
	"net/http"
)

//我们在文件中声明的 go:embed 指令和实际程序目录中的静态资源会以 IR 的方式产生关联，可以简单理解为此刻我们根据 go:embed 指令上下文中的变量已经被赋值了。
//go:embed assets
var assets embed.FS

func main() {
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(assets)))
	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
