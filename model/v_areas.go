package model

import (
	"time"

	"gorm.io/gorm"
)

type VArea struct {
	Version   string         `gorm:"column:version" json:"version"`
	LayerName string         `gorm:"column:layer_name" json:"layer_name"`
	StartTime time.Time      `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time      `gorm:"column:end_time" json:"end_time"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at"`
}
