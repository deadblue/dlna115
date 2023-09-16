package storage

// StorageService defines a interface which should be impled by storage service.
// It will be used by ContentDirectory service to browse files.
type StorageService interface {
	// Browse get items under specific parent folder
	Browse(parentId string) (items []Item)
	// TODO: More methods will be supported laaaaater ...

	// Fetch fetches file content from storage.
	Fetch(path string, offset int64, length int64) (content *Content, err error)
}
