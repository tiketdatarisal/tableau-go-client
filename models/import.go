package models

import "encoding/xml"

type Import struct {
	XMLName          xml.Name `xml:"import"`
	Import           string   `xml:"source,attr,omitempty"`
	DomainName       string   `xml:"domainName,attr,omitempty"`
	SiteRole         string   `xml:"siteRole,attr,omitempty"`
	GrantLicenseMode string   `xml:"grantLicenseMode,attr,omitempty"`
}
