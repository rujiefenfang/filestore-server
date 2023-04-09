package vo

import "time"

type UserInfoVO struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
