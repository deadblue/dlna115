package contentdirectory

import (
	"github.com/deadblue/dlna115/internal/mediaserver/service/storageservice"
)

type Service struct {
	ss storageservice.StorageService
}

func New(ss storageservice.StorageService) *Service {
	return &Service{
		ss: ss,
	}
}
