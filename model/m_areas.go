package model

import (
	"time"

	"gorm.io/gorm"
)

type GeometricType string

const (
	Point              GeometricType = "Point"
	Polygon            GeometricType = "Polygon"
	PolygonWithHole    GeometricType = "PolygonWithHole"
	Collection         GeometricType = "Collection"
	Linestring         GeometricType = "Linestring"
	MultiPoint         GeometricType = "MultiPoint"
	MultiLineString    GeometricType = "MultiLineString"
	MultiPolygon       GeometricType = "MultiPolygon"
	GeometryCollection GeometricType = "GeometryCollection"
)

type MArea struct {
	ID            uint           `gorm:"primarykey;column:ogc_fid" json:"id"`
	Geometry      string         `gorm:"column:wkb_geometry;type:geometry" json:"geometry"`
	GeometricType GeometricType  `gorm:"column:geometric_type" json:"geometric_type"`
	AreaCode      string         `gorm:"column:area_code" json:"area_code"`
	Name          string         `gorm:"column:name" json:"name"`
	Tags          string         `gorm:"column:tags" json:"tags"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at"`
}

func (MArea) TableName() string {
	return "public.m_areas"
}
