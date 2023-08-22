package didl

import "encoding/xml"

type Container interface {
	isContainer()
}

type Item interface {
	isItem()
}

type Document struct {
	XMLName       xml.Name `xml:"DIDL-Lite"`
	Namespace     string   `xml:"xmlns,attr"`
	NamespaceDc   string   `xml:"xmlns:dc,attr"`
	NamespaceUpnp string   `xml:"xmlns:upnp,attr"`
	NamespaceSec  string   `xml:"xmlns:sec,attr"`
	NamespaceDlna string   `xml:"xmlns:dlna,attr"`

	// Container objects
	Container []Container `xml:"container"`
	// Item objects
	Item []Item `xml:"item"`
}

func (d *Document) Init() *Document {
	d.Namespace = namespace
	d.NamespaceDc = namespaceDc
	d.NamespaceDlna = namespaceDlna
	d.NamespaceUpnp = namespaceUpnp
	return d
}

func (d *Document) AppendItem(item Item) {
	d.Item = append(d.Item, item)
}

func (d *Document) AppendContainer(container Container) {
	d.Container = append(d.Container, container)
}
