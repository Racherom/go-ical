package ical

import (
	"io"
	"net/url"
	"time"
)

// Encoder ...
type Encoder struct {
	w io.Writer
}

// VEvent ...
type VEvent struct {
	UID                              string
	DTSTAMP, CREATED, DTSTART, DTEND *time.Time
	Summary                          string
	Description                      string
	Location                         string
	URL                              *url.URL
}

// VCalendar ...
type VCalendar struct {
	URL    string
	Events []VEvent
}
