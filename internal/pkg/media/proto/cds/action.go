package cds

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

	UpdateID       string `xml:"UpdateID"`
	TotalMatches   int    `xml:"TotalMatches"`
	NumberReturned int    `xml:"NumberReturned"`
	Result         string `xml:"Result"`
}

func (r *BrowseResp) Init() *BrowseResp {
	r.NamespaceU = namespace
	return r
}
