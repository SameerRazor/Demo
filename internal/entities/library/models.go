package library

type Library struct {
	Book_ID   int    `gorm:"not null Index:idx_lib_bookid(255);not null" json:"book_id"`
	Aisle     int    `json:"aisle" gorm:"uniqueIndex:idx_aisle_position_level;not null"`
	Level     int    `json:"level" gorm:"uniqueIndex:idx_aisle_position_level;not null"`
	Position  int    `json:"position" gorm:"uniqueIndex:idx_aisle_position_level;not null"`
	CreatedAt string `gorm:"create_time" json:"created_at"`
	UpdatedAt string `gorm:"update_time" json:"updated_at"`
	IsDeleted bool   `json:"is_deleted"`
}
