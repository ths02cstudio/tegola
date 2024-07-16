package model

import "time"

type MArea struct {
	ID           uint64    `json:"ogc_fid"`
	Geometry     string    `json:"wkb_geometry"`
	GeometryType string    `json:"geometric_type"`
	AreaCode     string    `json:"area_code"`
	Properties   string    `json:"properties"`
	Tags         string    `json:"other_tags"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
