package main

import (
	"net/http"
	"text/template"
)

// 建立一个处理器
func testTemplate(w http.ResponseWriter, r *http.Request) {

	//解析模板文件
	t, _ := template.ParseFiles("hello.html")
	//执行模板
	t.Execute(w, "Hello World!")
}

func main() {

	http.HandleFunc("/testTemplate", testTemplate)
	http.ListenAndServe(":8080", nil)
}
