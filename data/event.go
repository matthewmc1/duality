package data

import (
	"time"

	"github.com/google/uuid"
)

type Evergreen struct {
	Id          uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	Label       []string  `json:"label,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
	Details     string    `json:"details"`
}
