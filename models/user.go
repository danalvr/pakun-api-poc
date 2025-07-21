package models

import "time"

type OTPEntry struct {
	Code      string
	ExpiresAt time.Time
}

var OTPStore = map[string]OTPEntry{}
