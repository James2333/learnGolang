package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.HandleFunc("/zbc",ProxyHandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080",nil))
}

func ProxyHandler(w http.ResponseWriter,r *http.Request)  {
	//生成一个新的url地址
	u,err:=url.Parse("https://www.baidu.com")
	if err != nil {
		return
	}
	//把url地址写入 到这个request里面去
	proxy:=httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL=u
		},
	}
	//转发
	log.Println("转发成功！")
	proxy.ServeHTTP(w,r)
}