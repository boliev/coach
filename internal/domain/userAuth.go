package domain

import "time"

type UserAuth struct {
	Token     string
	ExpiresAt time.Time
}
