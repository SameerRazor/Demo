package authorModels

type Author struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Biography   string    `json:"biography"`
	Nationality string    `json:"nationality"`
	DOB         int       `json:"int"`
}