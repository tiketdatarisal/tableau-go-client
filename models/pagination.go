package models

import "encoding/xml"

type Pagination struct {
	XMLName        xml.Name `xml:"pagination"`
	PageNumber     int      `xml:"pageNumber,attr,omitempty"`
	PageSize       int      `xml:"pageSize,attr,omitempty"`
	TotalAvailable int      `xml:"totalAvailable,attr,omitempty"`
}
