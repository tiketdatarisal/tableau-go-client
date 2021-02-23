package models

import "encoding/xml"

type QuerySitesResponse struct {
	XMLName     xml.Name    `xml:"tsResponse"`
	Pagination  *Pagination `xml:"pagination,omitempty"`
	Sites       *[]Site     `xml:"sites>site"`
}
