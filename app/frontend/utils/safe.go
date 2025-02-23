package utils

import (
	"net/url"
	"strings"
)

var validHosts = []string{
	"localhost:8080",
}

func ValidateNext(next string) bool {
	urlObj, err := url.Parse(next)
	if err != nil {
		return false
	}
	if inArray(urlObj.Host, validHosts) {
		return true
	}
	return false
}

func inArray(str string, list []string) bool {
	for _, v := range list {
		if strings.EqualFold(v, str) {
			return true
		}
	}
	return false
}
