package utils

import "github.com/pkg/browser"

// Open 打开到浏览器
func Open(url string) error {
	return browser.OpenURL(url)
}
