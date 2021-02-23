package models

import "encoding/xml"

type User struct {
	XMLName            xml.Name `xml:"user"`
	ID                 string   `xml:"id,attr"`
	Name               string   `xml:"name,attr,omitempty"`
	SiteRole           SiteRole `xml:"siteRole,attr,omitempty"`
	LastLogin          string   `xml:"lastLogin,attr,omitempty"`
	ExternalAuthUserID string   `xml:"externalAuthUserId,attr,omitempty"`
	AuthSetting        string   `xml:"authSetting,attr,omitempty"`
	Language           string   `xml:"language,attr,omitempty"`
	Locale             string   `xml:"locale,attr,omitempty"`
}
