package library

type Library struct {
	ID   int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Aisle     int       `gorm:"not null" json:"aisle"`
	Level     int       `gorm:"not null" json:"level"`
	Position  int       `gorm:"not null" json:"position"`
}
