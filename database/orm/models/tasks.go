package models

type Tasks struct {
	ID             int32 `gorm:"primaryKey"`
	Status         string
	ProcessPercent int16
	Content        string
	EndTime        int64 `gorm:"not null"`
	Updated        int64 `gorm:"autoUpdateTime"`
	Created        int64 `gorm:"autoCreateTime"`
}

func (s *Tasks) TableName() string {
	return "tasks"
}
