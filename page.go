package ButterCMS

import (
	"time"
)

type PageStatus string

const (
	PageStatusDraft     PageStatus = "draft"
	PageStatusPublished PageStatus = "published"
	PageStatusScheduled PageStatus = "scheduled"
)

type Page struct {
	Slug      string                 `json:"slug"`
	PageType  string                 `json:"page_type"`
	Fields    map[string]interface{} `json:"fields"`
	Status    PageStatus             `json:"status"`
	Published *time.Time             `json:"published,omitempty"`
	Updated   *time.Time             `json:"updated,omitempty"`
	Scheduled *time.Time             `json:"scheduled,omitempty"`
}
