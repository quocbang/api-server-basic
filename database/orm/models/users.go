package models

import "github.com/lib/pq"

type Users struct {
	Email    string        `gorm:"primaryKey"`
	Password []byte        `gorm:"type:bytea;not null"`
	Roles    pq.Int64Array `gorm:"type:smallint[];default:'{}';not null"`
}

func (u *Users) TableName() string {
	return "users"
}
