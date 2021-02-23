package models

import "encoding/xml"

type SwitchSiteRequest struct {
	XMLName xml.Name `xml:"tsRequest"`
	Site    Site     `xml:"site"`
}
