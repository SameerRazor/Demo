package author

type Author struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	AuthorName  string `gorm:"not null;unique" json:"author_name"`
	Biography   string `json:"biography"`
	DateOfBirth string `json:"date_of_birth"`
	Nationality string `json:"nationality"`
	CreatedAt   string `gorm:"create_time"`
	UpdatedAt   string `gorm:"update_time"`
	IsDeleted   bool   `json:"is_deleted"`
}
