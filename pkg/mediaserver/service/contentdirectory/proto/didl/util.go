package didl

import "fmt"

func FormatDuration(seconds float64) string {
	m := int(seconds / 60)
	s := seconds - float64(m*60)
	h := int(m / 60)
	m = m % 60
	return fmt.Sprintf("%d:%02d:%06.3f", h, m, s)
}
