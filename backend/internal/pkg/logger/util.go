package logger

import (
	"strings"
)

func IsSkipLog(contentType string) bool {
	lists := []string{
		"application/tar+gzip",
		"application/octet-stream",
		"multipart/form-data",
	}

	for _, list := range lists {
		if strings.Contains(contentType, list) {
			return true
		}
	}

	return false
}
