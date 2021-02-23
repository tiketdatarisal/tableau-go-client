package models

import "encoding/xml"

type GetGroupsForUserResponse struct {
	XMLName    xml.Name    `xml:"tsResponse"`
	Pagination *Pagination `xml:"pagination,omitempty"`
	Groups     *[]Group    `xml:"groups>group"`
}
