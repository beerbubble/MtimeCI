package models

import (
	"encoding/xml"
)

type RunJobExecutions struct {
	XMLName xml.Name          `xml:"executions"`
	Count   int               `xml:"count,attr"`
	Exs     []RunJobExecution `xml:"execution"`
}

type RunJobExecution struct {
	XMLName xml.Name `xml:"execution"`
	Id      string   `xml:"id,attr"`
	Href    string   `xml:"href,attr"`
	Status  string   `xml:"status,attr"`
	Project string   `xml:"project,attr"`
	User    string   `xml:"user"`
}
