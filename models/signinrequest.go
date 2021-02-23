package models

import "encoding/xml"

type SignInRequest struct {
	XMLName     xml.Name    `xml:"tsRequest"`
	Credentials Credentials `xml:"credentials"`
}
