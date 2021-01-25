package main

import "time"

//Bulletin struct
type Bulletin struct {
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	Details         string    `json:"details"`
	CreateTimestamp time.Time `json:"date"`
}

//Ads struct
type Ads struct {
	Title           string    `json:"title" binding:"required"`
	Company         string    `json:"company" binding:"required"`
	Product         string    `json:"product" binding:"required"`
	CreateTimestamp time.Time `json:"date" binding:"required"`
}
