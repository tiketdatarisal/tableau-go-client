package models

import "encoding/xml"

type Group struct {
	XMLName xml.Name `xml:"group"`
	ID      string   `xml:"id,attr,omitempty"`
	Name    string   `xml:"name,attr,omitempty"`
	Domain  *Domain  `xml:"domain,omitempty"`
	Import  *Import  `xml:"import,omitempty"`
}
