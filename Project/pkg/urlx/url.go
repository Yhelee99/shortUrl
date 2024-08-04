package urlx

import (
	"errors"
	"net/url"
	"path"
)

// GetBasePath 获取最后一级路径
func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", err
	}

	if myUrl.Host == "" {
		return "", errors.New("no Host in target Url")
	}

	return path.Base(myUrl.Path), nil
}
