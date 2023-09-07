package storage

import "net/http"

// StorageService defines a interface which should be impled by storage service.
// It will be used by ContentDirectory service to browse files.
type StorageService interface {
	MountTo(mux *http.ServeMux)

	Browse(parentId string) (items []Item)
	// TODO: More methods will be supported laaaaater ...
}
