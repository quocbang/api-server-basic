package models

import (
	"gorm.io/gorm/schema"
)

type Models interface {
	schema.Tabler
}

func ListModels() []Models {
	return []Models{
		&Users{},
		&Tasks{},
	}
}
