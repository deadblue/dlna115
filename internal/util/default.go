package util

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func DefaultNumber[T Number](value T, defValue T) T {
	if value != 0 {
		return value
	} else {
		return defValue
	}
}

func DefaultString(value string, defValue string) string {
	if value != "" {
		return value
	} else {
		return defValue
	}
}

func DefaultStringFunc(value string, defFunc func() string) string {
	if value != "" {
		return value
	} else {
		return defFunc()
	}
}
