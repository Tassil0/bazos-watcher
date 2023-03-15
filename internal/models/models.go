package models

import "time"

type Query struct {
	Id              uint32 `gorm:"primaryKey" json:"id"`
	Query           string `gorm:"not null" json:"query"`
	FilterPriceLow  uint32 `json:"filterPriceLow"`
	FilterPriceHigh uint32 `json:"filterPriceHigh"`
	FilterCity      string `json:"filterCity"`
}

type Item struct {
	Id          uint32    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`
	Price       uint32    `gorm:"not null" json:"price"`
	Img         string    `gorm:"not null" json:"img"`
	Location    string    `gorm:"not null" json:"'location'"`
	Url         string    `gorm:"not null" json:"url"`
	Blacklisted bool      `gorm:"not null" json:"blacklisted"`
	DateAdded   time.Time `gorm:"not null" json:"dateAdded"`
	LastUpdated time.Time `gorm:"not null" json:"lastUpdated"`
	QueryId     uint32    `gorm:"not null;unique" json:"queryId"`
}
