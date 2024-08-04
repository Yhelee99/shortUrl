package main

import (
	"fmt"
	"log"
	"net/url"
	"path"
)

func main() {
	targetUrl := "https://google.com/foo/bar?a=1&b=2"
	myUrl, err := url.Parse(targetUrl)
	/*
		url.Parse()出来
		Scheme:"http"	     协议
		Host:"google.com"    域名
		Path:"/foo/bar"		 路径
		RawQuery:"a=1&b=2"	 查询参数
	*/
	if err != nil {
		log.Fatalf("Parse failed,err:%v\n", err)
	}
	fmt.Printf("myurl:%#v\n", myUrl)

	// path.Base() 获取最后一级路径
	fmt.Println(path.Base(myUrl.Path))

}
