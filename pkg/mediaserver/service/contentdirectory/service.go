package contentdirectory

import "github.com/deadblue/dlna115/pkg/storage"

type Service struct {
	ss storage.StorageService
}

func New(ss storage.StorageService) *Service {
	return &Service{
		ss: ss,
	}
}
