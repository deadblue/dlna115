package proto

import "encoding/xml"

type BrowseReq struct {
	XMLName    xml.Name
	NamespaceU string `xml:"xmlns:u,attr"`

	ObjectID       string `xml:"ObjectID"`
	BrowseFlag     string `xml:"BrowseFlag"`
	Filter         string `xml:"Filter"`
	StartingIndex  int    `xml:"StartingIndex"`
	RequestedCount int    `xml:"RequestedCount"`
	SortCriteria   string `xml:"SortCriteria"`
}

type BrowseResp struct {
	XMLName    xml.Name `xml:"u:BrowseResponse"`
	NamespaceU string   `xml:"xmlns:u,attr"`

	UpdateID       uint   `xml:"UpdateID"`
	TotalMatches   int    `xml:"TotalMatches"`
	NumberReturned int    `xml:"NumberReturned"`
	Result         string `xml:"Result"`
}

func (r *BrowseResp) Init() *BrowseResp {
	r.NamespaceU = namespace
	return r
}

type SearchReq struct {
	XMLName    xml.Name
	NamespaceU string `xml:"xmlns:u,attr"`

	ContainerID    string `xml:"ContainerID"`
	SearchCriteria string `xml:"SearchCriteria"`
	Filter         string `xml:"Filter"`
	StartingIndex  int    `xml:"StartingIndex"`
	RequestedCount int    `xml:"RequestedCount"`
	SortCriteria   string `xml:"SortCriteria"`
}

type SearchResp struct {
	XMLName    xml.Name `xml:"u:BrowseResponse"`
	NamespaceU string   `xml:"xmlns:u,attr"`

	UpdateID       string `xml:"UpdateID"`
	TotalMatches   int    `xml:"TotalMatches"`
	NumberReturned int    `xml:"NumberReturned"`
	Result         string `xml:"Result"`
}

func (r *SearchResp) Init() *SearchResp {
	r.NamespaceU = namespace
	return r
}
