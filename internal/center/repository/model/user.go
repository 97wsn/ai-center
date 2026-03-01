package model

import (
	"time"
)

type User struct {
	Id        int64     `json:"id" db:"id"`                 // PK
	CreatedAt time.Time `json:"created_at" db:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"` // 更新时间
}

func (User) TableName() string {
	return "user"
}

func (User) PK() string {
	return "id"
}
