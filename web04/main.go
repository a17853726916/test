package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID       int
	Username string
	Password string
}

func handler(w http.ResponseWriter, r *http.Request) {
	//设置响应头中内容的类型
	w.Header().Set("Content-Type", "application/json")
	user := User{
		ID:       1,
		Username: "admin",
		Password: "123456",
	}
	//将 user 转换为 json 格式
	json, _ := json.Marshal(user)
	w.Write(json)
}
func main() {
	http.HandleFunc("/testJson", handler)
	http.ListenAndServe(":8080",nil)
}
