package main

import (
	"embed"
	"log"
	"net/http"
	"net/http/pprof"
	"runtime"
)

//我们在文件中声明的 go:embed 指令和实际程序目录中的静态资源会以 IR 的方式产生关联，可以简单理解为此刻我们根据 go:embed 指令上下文中的变量已经被赋值了。
//go:embed assets
var assets embed.FS

func registerRoute() *http.ServeMux {
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(assets)))
	return mutex
}

func enableProf(mutex *http.ServeMux) {
	runtime.GOMAXPROCS(2)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	mutex.HandleFunc("/debug/pprof/", pprof.Index)
	mutex.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mutex.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mutex.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mutex.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

func main() {
	mutex := registerRoute()
	enableProf(mutex)

	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
