package urlx

import (
	"net/url"
	"path"
)

// GetBasePath 获取最后一级路径
func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", err
	}
	return path.Base(myUrl.Path), nil
}
