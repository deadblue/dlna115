package didl

const (
	ItemClassAudio          = "object.item.audioItem"
	ItemClassMusicTrack     = "object.item.audioItem.musicTrack"
	ItemClassAudioBroadcast = "object.item.audioItem.audioBroadcast"
	ItemClassAudioBook      = "object.item.audioItem.audioBook"

	ItemClassVideo          = "object.item.videoItem"
	ItemClassMovie          = "object.item.videoItem.movie"
	ItemClassVideoBroadcast = "object.item.videoItem.videoBroadcast"
	ItemClassMusicVideoClip = "object.item.videoItem.musicVideoClip"

	ItemClassImage = "object.item.imageItem"
	ItemClassPhoto = "object.item.imageItem.photo"

	ItemClassPlaylistItem = "object.item.playlistItem"

	ItemClassText = "object.item.textItem"
)

type _BaseItem struct {
	Object
	Res Res `xml:"res"`
}

func (o *_BaseItem) isItem() {}

type ImageItem struct {
	_BaseItem
}

func (o *ImageItem) Init() *ImageItem {
	o.Class = ItemClassImage
	o.Restricted = "1"
	return o
}

type AudioItem struct {
	_BaseItem
}

func (o *AudioItem) Init() *AudioItem {
	o.Class = ItemClassAudio
	o.Restricted = "1"
	return o
}

type VideoItem struct {
	// Derived from _BaseItem
	_BaseItem
}

func (o *VideoItem) Init() *VideoItem {
	o.Class = ItemClassVideo
	o.Restricted = "1"
	return o
}
