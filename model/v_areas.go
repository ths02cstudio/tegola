package model

import "time"

type VArea struct {
	Version   string    `json:"version"`
	LayerName string    `json:"layer_name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
