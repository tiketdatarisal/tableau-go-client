package models

import "encoding/xml"

type AddUserToGroupResponse struct {
	XMLName xml.Name `xml:"tsResponse"`
	User    *User    `xml:"user"`
}
