package model

type Location struct {
	ID int `gorm:"primaryKey"`
	InputText string
	QueryText string
	Valid bool
	Lat float64
	Lon float64
	Name string
	Confidence string
	EntityType string
	MatchCode string
}

