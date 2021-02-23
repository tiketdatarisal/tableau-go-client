package models

import "encoding/xml"

type QueryGroupsResponse struct {
	XMLName    xml.Name    `xml:"tsResponse"`
	Pagination *Pagination `xml:"pagination,omitempty"`
	Groups     *[]Group    `xml:"groups>group"`
}
