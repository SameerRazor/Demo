package author

type Author struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	AuthorName  string `gorm:"not null" json:"author_name"`
	Biography   string `json:"biography"`
	DateOfBirth int    `json:"date_of_birth"`
	Nationality string `json:"nationality"`
}
