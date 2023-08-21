package forward

import "fmt"

func (s *Service) GetAccessURL(accessCode string) string {
	return fmt.Sprintf("%s%s", HandleURL, accessCode)
}
