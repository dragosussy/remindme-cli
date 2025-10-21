package model

import "time"

type Reminder struct {
	Id             string
	Title          string
	Text           string
	CronExpression string
	NextRunAt      time.Time
	Acknowledged   bool
}
