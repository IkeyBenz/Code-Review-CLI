package models

import "time"

// Request struct serves as a model for all request objects in our database
type Request struct {
	_id           string
	from          User
	to            User
	crRequest     string
	crResponse    string
	dateRequested *time.Time
	dateResponded *time.Time
	opened        bool
}
