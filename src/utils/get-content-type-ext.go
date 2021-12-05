package utils

import (
	"strings"
)

func GetContentTypeExt(contentType string) string {
	trimmed := strings.TrimSpace(contentType)

	switch trimmed {
	case "image/png":
		return ".png"
	case "image/jpeg":
		return ".jpg"
	case "image/webp":
		return "webp"
	}

	panic("unsupported content type (" + trimmed + ")")
}
