package models

import "time"

type Weather struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	Lon       float64   `json:"lon"`
	Lat       float64   `json:"lat"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    int64     `json:"userId"`
}
