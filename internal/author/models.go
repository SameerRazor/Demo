package author

type Author struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	AuthorName  string `gorm:"not null" json:"authorname"`
	Biography   string `json:"biography"`
	DateOfBirth int    `json:"dateofbirth"`
	Nationality string `json:"nationality"`
}
