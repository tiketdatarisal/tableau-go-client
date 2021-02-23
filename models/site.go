package models

import "encoding/xml"

type Site struct {
	XMLName                      xml.Name `xml:"site"`
	ID                           string   `xml:"id,attr,omitempty"`
	ContentURL                   string   `xml:"contentUrl,attr"`
	AdminMode                    string   `xml:"adminMode,attr,omitempty"`
	DisableSubscription          bool     `xml:"disableSubscription,attr,omitempty"`
	UserQuota                    int64    `xml:"userQuota,attr,omitempty"`
	StorageQuota                 int64    `xml:"storageQuota,attr,omitempty"`
	State                        string   `xml:"state,attr,omitempty"`
	StatusReason                 string   `xml:"statusReason,attr,omitempty"`
	RevisionHistoryEnabled       bool     `xml:"revisionHistoryEnabled,attr,omitempty"`
	RevisionLimit                int64    `xml:"revisionLimit,attr,omitempty"`
	SubscribeOthersEnabled       bool     `xml:"subscribeOthersEnabled,attr,omitempty"`
	AllowSubscriptionAttachments bool     `xml:"allowSubscriptionAttachments,attr,omitempty"`
	GuestAccessEnabled           bool     `xml:"guestAccessEnabled,attr,omitempty"`
	CacheWarmupEnabled           bool     `xml:"cacheWarmupEnabled,attr,omitempty"`
	CommentingEnabled            bool     `xml:"commentingEnabled,attr,omitempty"`
	EditingFlowsEnabled          bool     `xml:"editingFlowsEnabled,attr,omitempty"`
	SchedulingFlowsEnabled       bool     `xml:"schedulingFlowsEnabled,attr,omitempty"`
	ExtractEncryptionMode        string   `xml:"extractEncryptionMode,attr,omitempty"`
	CatalogingEnabled            bool     `xml:"catalogingEnabled,attr,omitempty"`
	DerivedPermissionEnabled     bool     `xml:"derivedPermissionEnabled,attr,omitempty"`
	RequestAccessEnabled         bool     `xml:"requestAccessEnabled,attr,omitempty"`
	RunNowEnabled                bool     `xml:"runNowEnabled,attr,omitempty"`
	IsDataAlertsEnabled          bool     `xml:"isDataAlertsEnabled,attr,omitempty"`
	AskDataMode                  string   `xml:"askDataMode,attr,omitempty"`
	UseDefaultTimeZone           bool     `xml:"useDefaultTimeZone,attr,omitempty"`
	TimeZone                     string   `xml:"timeZone,attr,omitempty"`
}
