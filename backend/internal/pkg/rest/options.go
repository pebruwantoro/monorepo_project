package rest

import (
	"time"
)

type Options struct {
	ClientID string
	Address  string
	Timeout  time.Duration
	SkipTLS  bool
}
