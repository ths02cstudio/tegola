package model

import "time"

type DArea struct {
	ID            uint64    `json:"geom_id"`
	OgcFid        uint64    `json:"ogc_fid"`
	ExtProperties string    `json:"ext_properties"`
	ExtTags       string    `json:"ext_tags"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
