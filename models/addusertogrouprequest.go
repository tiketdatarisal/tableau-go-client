package models

import "encoding/xml"

type AddUserToGroupRequest struct {
	XMLName xml.Name `xml:"tsRequest"`
	User    User     `xml:"user"`
}
