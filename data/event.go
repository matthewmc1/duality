package data

import (
	"github.com/google/uuid"
)

type Evergreen struct {
	Id          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Label       string    `db:"labels"`
	CreatedDate string    `db:"created_date"`
	Details     string    `dvb:"details"`
}
