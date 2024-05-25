package entity

import "time"

type Organization struct {
	ID        uint
	Name      string
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *Organization) TableName() string {
	return "organizations"
}
