package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过详细配置处理器处理请求")
}

func main() {

	myHandler := MyHandler{}

	// 调用处理器
	// http.Handle("/myHandler", &myHandler)
	server := http.Server{
		Addr:        "8080",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()

}
