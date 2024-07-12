package main

import (
	"fmt"
	"net/http"
)

// 设置 CORS 头
func setCORSHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许所有域名
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
}

// 处理预检请求
func handlePreflight(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == "OPTIONS" {
		setCORSHeaders(w, r)
		w.WriteHeader(http.StatusOK)
		return true
	}
	return false
}

// 主处理器
func handler(w http.ResponseWriter, r *http.Request) {
	//处理预检请求
	if handlePreflight(w, r) {
		return
	}

	//设置CORS头
	setCORSHeaders(w, r)

	//处理实际请求
	w.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}
