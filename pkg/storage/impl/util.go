package impl

import (
	"mime"
	"path/filepath"
	"strings"
)

var (
	_ImageExt = map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".gif":  true,
	}
)

func getMimeType(fileName string) string {
	ext := filepath.Ext(fileName)
	return mime.TypeByExtension(ext)
}

func isImageFile(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	return _ImageExt[ext]
}
