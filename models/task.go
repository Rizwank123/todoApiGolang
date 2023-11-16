package models

import (
	"time"

)

type Task struct {
	ID        int       `db:"id" json:"id,omitempty"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc,omitempty"`
	CreatedAt time.Time ` json:"created_at"`
	UserId    int       `db:"user" json:"user_id,omitempty"`
	Status    string    `json:"status"`
}
