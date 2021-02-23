package models

import "encoding/xml"

type QueryUserOnSiteResponse struct {
	XMLName xml.Name `xml:"tsResponse"`
	User    *User    `xml:"user,omitempty"`
	Domain  *Domain  `xml:"domain,omitempty"`
}
