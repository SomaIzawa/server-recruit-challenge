package model

type AlbumID int

type Album struct {
	ID       AlbumID  `json:"id" gorm:"primaryKey"`
	Title    string   `json:"title"`
	Singer Singer `gorm:"foreignKey:SingerID; constraint:OnDelete:CASCADE"`
	SingerID SingerID `json:"singer_id" gorm:"not_null"` // モデル Singer の ID と紐づきます
}
