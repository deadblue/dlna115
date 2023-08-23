package storageservice

// StorageService defines a interface which should be impled by storage service.
// It will be used by ContentDirectory service to browse files.
type StorageService interface {
	Browse(parentId string) (items []Item)
	// TODO: More methods will be supported laaaaater ...
}
