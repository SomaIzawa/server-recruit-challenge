package model

type SingerID int

type Singer struct {
	ID   SingerID `json:"id" gorm:"primaryKey"`
	Name string   `json:"name"`
}
