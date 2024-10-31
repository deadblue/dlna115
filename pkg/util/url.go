package util

import (
	"fmt"
	"net/url"
	"strings"
)

func GetRootUrl(rawUrl string) string {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host)
}

func GetBaseUrl(rawUrl string) string {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	path := u.EscapedPath()
	if slashIndex := strings.LastIndex(u.Path, "/"); slashIndex > 0 {
		path = path[:slashIndex+1]
	}
	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, path)
}

func isFullUrl(s string) bool {
	index := strings.Index(s, "://")
	return index > 0
}

func GetAbsoluteUrl(originUrl, relativePath string) string {
	if isFullUrl(relativePath) {
		return relativePath
	}
	if strings.HasPrefix(relativePath, "/") {
		return GetRootUrl(originUrl) + relativePath
	} else {
		return GetBaseUrl(originUrl) + relativePath
	}
}
