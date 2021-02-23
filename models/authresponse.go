package models

import "encoding/xml"

type AuthResponse struct {
	XMLName     xml.Name    `xml:"tsResponse"`
	Credentials Credentials `xml:"credentials"`
}
