package didl

const (
	ContainerClassAlbum      = "object.container.alnum"
	ContainerClassMusicAlbum = "object.container.alnum.musicAlbum"
	ContainerClassPhotoAlbum = "object.container.alnum.photoAlbum"

	ContainerClassGenre      = "object.container.genre"
	ContainerClassMusicGenre = "object.container.genre.musicGenre"
	ContainerClassMovieGenre = "object.container.genre.movieGenre"

	ContainerClassPlaylist = "object.container.playlistContainer"

	ContainerClassPerson      = "object.container.person"
	ContainerClassMusicArtist = "object.container.person.musicArtist"

	ContainerClassStorageSystem = "object.container.storageSystem"
	ContainerClassStorageVolume = "object.container.storageVolume"
	ContainerClassStorageFolder = "object.container.storageFolder"
)

type _BaseContainer struct {
	Object
}

func (o *_BaseContainer) isContainer() {}

// Ref: 7.11 @ UPnP-av-ContentDirectory-v1-Service.pdf
type StorageVolumeContainer struct {
	_BaseContainer
	StorageTotal  int64  `xml:"upnp:storageTotal"`
	StorageUsed   int64  `xml:"upnp:storageUsed"`
	StorageFree   int64  `xml:"upnp:storageFree"`
	StorageMedium string `xml:"upnp:storageMedium"`
}

func (o *StorageVolumeContainer) Init() *StorageVolumeContainer {
	o.Restricted = "1"
	o.Class = ContainerClassStorageVolume
	return o
}

// Ref: 7.12 @ UPnP-av-ContentDirectory-v1-Service.pdf
type StorageFolderContainer struct {
	_BaseContainer
	StorageUsed int `xml:"upnp:storageUsed"`
}

func (o *StorageFolderContainer) Init() *StorageFolderContainer {
	o.Restricted = "1"
	o.Class = ContainerClassStorageFolder
	return o
}
