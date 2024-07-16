package model

import (
	"time"

	"gorm.io/gorm"
)

type DArea struct {
	ID            uint64         `gorm:"column:geom_id" json:"geom_id"`
	OgcFid        uint           `gorm:"column:ogc_fid" json:"ogc_fid"`
	ExtProperties string         `gorm:"column:ext_properties" json:"ext_properties"`
	ExtTags       string         `gorm:"column:ext_tags" json:"ext_tags"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at"`
}

func (DArea) TableName() string {
	return "public.d_areas"
}
