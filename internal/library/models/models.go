package library

type Library struct {
	ID       int `gorm:"primaryKey" json:"book_id"`
	Aisle    int `json:"aisle"`
	Level    int `json:"level"`
	Position int `json:"position"`
}
