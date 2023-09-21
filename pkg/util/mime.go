package util

import (
	"strings"
)

const (
	MimeTypeM3U8 = "application/vnd.apple.mpegurl"

	_DefaultMimeType = "application/octet-stream"
)

var (
	_MimeTypes = map[string]string{
		// Audio
		"aac":  "audio/aac",
		"ac3":  "audio/ac3",
		"dts":  "audio/vnd.dts",
		"flac": "audio/x-flac",
		"m4a":  "audio/x-m4a",
		"mka":  "audio/x-matroska",
		"mp3":  "audio/mpeg",
		"oga":  "audio/ogg",
		"ogg":  "audio/ogg",
		"ra":   "audio/x-pn-realaudio",
		"ram":  "audio/x-pn-realaudio",
		"wav":  "audio/x-wav",
		"wma":  "audio/x-ms-wma",
		// Audio (non-standard)
		"ape": "audio/x-ape",
		"tak": "audio/x-tak",

		// Image
		"bmp":  "image/x-ms-bmp",
		"gif":  "image/gif",
		"jpeg": "image/jpeg",
		"jpg":  "image/jpeg",
		"ico":  "image/x-icon",
		"png":  "image/png",
		"svg":  "image/svg+xml",
		"tif":  "image/tiff",
		"tiff": "image/tiff",
		"webp": "image/webp",

		// Video
		"asf":  "video/x-ms-asf",
		"avi":  "video/x-msvideo",
		"f4v":  "video/x-f4v",
		"flv":  "video/x-flv",
		"m2ts": "video/mp2t",
		"m4v":  "video/mp4",
		"mkv":  "video/x-matroska",
		"mov":  "video/quicktime",
		"mp4":  "video/mp4",
		"mpeg": "video/mpeg",
		"mpg":  "video/mpeg",
		"ogv":  "video/ogg",
		"rm":   "application/vnd.rn-realmedia",
		"rmvb": "application/vnd.rn-realmedia-vbr",
		"ts":   "video/mp2t",
		"webm": "video/webm",
		"wmv":  "video/x-ms-wmv",
		// Video (non-standard)
		"divx": "video/x-divx",
	}
)

func GetMimeTypeForExt(ext string) string {
	ext = strings.ToLower(ext)
	if mt, ok := _MimeTypes[ext]; ok {
		return mt
	} else {
		return _DefaultMimeType
	}
}

func GetMimeType(filename string) string {
	dotIndex := strings.LastIndex(filename, ".")
	if dotIndex >= 0 {
		return GetMimeTypeForExt(filename[dotIndex+1:])
	} else {
		return _DefaultMimeType
	}
}

func IsImageFile(filename string) bool {
	return strings.HasPrefix(
		GetMimeType(filename), "image/",
	)
}
