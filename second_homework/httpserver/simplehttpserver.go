package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/log", logHandler)
	http.HandleFunc("/healthz", healthzHandler)
	err := http.ListenAndServe(":8099", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

// client定义header，server接收request中的header并添加到response的header中
/* 使用curl模拟请求
> curl.exe  127.0.0.1:8085/header -H 'User-Agent: Chrome 94' -I
HTTP/1.1 403 Forbidden
Server: Nginx-1.12.1
User-Agent: Chrome 94
Date: Wed, 29 Sep 2021 07:40:06 GMT
Content-Length: 10
Content-Type: text/plain; charset=utf-8
*/
func headerHandler(w http.ResponseWriter, r *http.Request) {
	////接收客户端UA
	//ua := r.UserAgent()
	////自定义header
	//w.Header().Set("Server", "Nginx-1.12.1")
	////把rquest的UA返回到response的header中
	//w.Header().Set("User-Agent", ua)
	////自定义网页返回状态码
	//// w.WriteHeader(403)
	////回显客户端消息
	//w.Write([]byte("add_header"))

	//打印请求方法，请求URL,请求协议
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	//打印请求头
	fmt.Fprintln(w, "打印请求头：")
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	//打印请求的主机地址
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	//打印与远端地址
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	//打印请求体
	fmt.Fprintln(w, "打印请求体：")
	for k, v := range r.Form {
		fmt.Fprintf(w, "From[%q] = %q\n", k, v)
	}
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	go_version := runtime.Version()
	w.Header().Set("GoVersion", go_version)
	w.Write([]byte("GoVersion: " + go_version))
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	//日志格式：远端地址-请求方法-请求地址-请求头-请求体
	fmt.Print(r.RemoteAddr, "-")     //打印客户端地址
	fmt.Print("-", r.Method)         //打印请求方法
	fmt.Print("-", r.URL)            //打印访问的url
	fmt.Print("-", r.Header)         //打印请求头
	fmt.Print("-", r.Body)           //打印请求体
	w.Write([]byte("Hello Golang!")) //回显客户端消息
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)     //设置状态码200
	w.Write([]byte("200")) //客户端显示200
}
