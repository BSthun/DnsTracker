package models

import "time"

type Log struct {
	Time     time.Time `json:"time"`
	Hostname string    `json:"hostname"`
	Status   int       `json:"status"`
}
