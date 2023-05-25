package models

type Genre struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Genre     string    `gorm:"not null" json:"genre"`
}
