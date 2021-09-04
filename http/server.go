package http

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Entry() {
	// 路由
	http.HandleFunc("/", AppHome)

	addr := ":9090"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("服务器启动时遇到些问题...", err)
	}
}

func AppHome(w http.ResponseWriter, req *http.Request) {
	log.Printf("请求路径: %v\n", req.URL.Path)

	err := req.ParseForm()
	if err != nil {
		log.Fatal("表单解析失败", err)
		return
	}
	log.Printf("请求参数表单: %v\n", req.Form)

	if len(req.Form["name"]) > 0 {
		log.Printf("获取name请求参数: %v\n", req.Form["name"][0])
	}

	if len(req.Form) > 0 {
		log.Println("开始解析请求参数...")
		for k, v := range req.Form {
			log.Printf("k: %v, val: %v", k, strings.Join(v, ","))
		}
		log.Printf("请求参数解析完毕! 共%v个\n\n", len(req.Form))
	}

	_, err = fmt.Fprintf(w, "Hello Web!")
	if err != nil {
		log.Fatal("响应失败", err)
		return
	}
}
