package main

import (
	"fmt"
	"net/http"
)

// 建立一个处理器
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "已经发送请求地址", r.URL.Path)
	fmt.Fprintln(w, "已经发送请求地址后请求的查询字符串", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中所有的信息", r.Header)
	fmt.Fprintln(w, "请求头中所有的信息“Accept-Encoding”", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中所有的信息“Accept-Encoding”的属性值", r.Header.Get("Accept-Encoding"))

	// // 获取请求体中内容的长度
	// len := r.ContentLength
	// // 创建byte切片
	// body := make([]byte, len)
	// // 将请求体中的读取到body中
	// r.Body.Read(body)
	// // 在浏览器中显示
	// fmt.Fprintln(w, "请求体中的内容有", string(body))

	// 解析表单 在调用r.Form之前必须执行该操作
	r.ParseForm()
	// fmt.Println(err)
	// 获取请求参数
	fmt.Fprintln(w, "请求参数有", r.Form["username"])
	fmt.Fprintln(w, "Post请求的Form表单参数有", r.PostForm)

	fmt.Fprintln(w, "请求参数username的值为：", r.FormValue("username"))

}

func main() {

	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
