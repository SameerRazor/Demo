package library

type Library struct {
	ID       int `gorm:"primaryKey" json:"id"`
	Aisle    int `json:"aisle"`
	Level    int `json:"level"`
	Position int `json:"position"`
}
