package models

import "encoding/xml"

type Credentials struct {
	XMLName   xml.Name `xml:"credentials"`
	Token     string   `xml:"token,attr,omitempty"`
	PATName   string   `xml:"personalAccessTokenUser,attr,omitempty"`
	PATSecret string   `xml:"personalAccessTokenSecret,attr,omitempty"`
	Name      string   `xml:"name,attr,omitempty"`
	Password  string   `xml:"password,attr,omitempty"`
	Site      *Site    `xml:"site,omitempty"`
	User      *User    `xml:"user,omitempty"`
}
