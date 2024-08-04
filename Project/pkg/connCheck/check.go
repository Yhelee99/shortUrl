package conncheck

import (
	"net/http"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true, // 关闭HTTP keep-alive 不需要保持链接
	},
	Timeout: 2 * time.Second,
}

func CheckUrl(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	return resp.StatusCode == http.StatusOK // 如果请求的是带跳转的链接，此处也不予放行
}
