package models

import "encoding/xml"

type GetUsersOnSiteResponse struct {
	XMLName    xml.Name    `xml:"tsResponse"`
	Pagination *Pagination `xml:"pagination,omitempty"`
	Users      *[]User     `xml:"users>user"`
}
