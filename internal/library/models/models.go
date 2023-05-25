package libraryModels

type Library struct {
	Book_ID   int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Aisle     int       `gorm:"not null" json:"aisle"`
	Level     int       `gorm:"not null" json:"level"`
	Position  int       `gorm:"not null" json:"position"`
}
