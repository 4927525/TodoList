package serializer

import (
	"time"
)

type BuildUser struct {
	Uid       uint      `json:"uid"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}