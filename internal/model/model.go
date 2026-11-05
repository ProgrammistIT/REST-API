package model

import (
	"time"
)

type Threat struct {
	MD5          string    `json:"md5"`
	Description  string    `json:"description"`
	Name         string    `json:"name"`
	Severity     string    `json:"severity"`
	DetectedTime time.Time `json:"detected_time"`
}
