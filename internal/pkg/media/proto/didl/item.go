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
}

func (o *_BaseItem) isItem() {}

type VideoItem struct {
	// Derived from _BaseItem
	_BaseItem
	// Video item should has Res
	Res Res `xml:"res"`
}

func (o *VideoItem) Init() *VideoItem {
	o.Class = ItemClassVideo
	return o
}
