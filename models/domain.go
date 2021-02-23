package models

import "encoding/xml"

type Domain struct {
	XMLName xml.Name `xml:"domain"`
	Name    string   `xml:"name,attr"`
}
