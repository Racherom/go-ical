package ical

import (
	"fmt"
	"io"
	"strings"
)

const timeFormat = "20060102T150405Z"

// NewEncoder ...
func NewEncoder(w io.Writer) *Encoder {

	return &Encoder{w: w}
}

// Encode ...
func (ec *Encoder) Encode(cal VCalendar) {
	fmt.Fprintln(ec.w, "BEGIN:VCALENDAR\nVERSION:2.0\nMETHOD:PUBLISH")

	for _, e := range cal.Events {
		fmt.Fprintf(ec.w, "BEGIN:VEVENT\nUID:%s@%s\nCLASS:PUBLIC\n", e.UID, cal.URL)

		if e.Location != "" {
			fmt.Fprintf(ec.w, "LOCATION:%s\n", strings.ReplaceAll(e.Location, "\n", `\n`))
		}

		if e.Summary != "" {
			fmt.Fprintf(ec.w, "SUMMARY:%s\n", strings.ReplaceAll(e.Summary, "\n", `\n`))
		}

		if e.Description != "" {
			fmt.Fprintf(ec.w, "DESCRIPTION:%s\n", strings.ReplaceAll(e.Description, "\n", `\n`))
		}

		if e.DTSTAMP != nil {
			fmt.Fprintf(ec.w, "DTSTAMP:%s\n", e.DTSTAMP.Format(timeFormat))
		}

		if e.DTSTART != nil {
			fmt.Fprintf(ec.w, "DTSTART:%s\n", e.DTSTART.Format(timeFormat))
		}

		if e.DTEND != nil {
			fmt.Fprintf(ec.w, "DTEND:%s\n", e.DTEND.Format(timeFormat))
		}

		if e.CREATED != nil {
			fmt.Fprintf(ec.w, "CREATED:%s\n", e.CREATED.Format(timeFormat))
		}

		if e.URL != nil {
			fmt.Fprintf(ec.w, "URL:%s\n", e.URL.String())
		}

		fmt.Fprintln(ec.w, "END:VEVENT")
	}
	fmt.Fprintln(ec.w, "END:VCALENDAR")
}
