package entities

import "time"

type Player struct {
	Name         string    `json:"name" binding:"required"`
	Age          int       `json:"age" binding:"required"`
	CreationTime time.Time `json:"creation_time"`
}
