package util

var (
	osVer = getOsVersion()
)

func OsVersion() string {
	return osVer
}
