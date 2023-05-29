package genre

type Genre struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Genre     string `gorm:"not null" json:"genre"`
	CreatedAt string `gorm:"create_time"`
	UpdatedAt string `gorm:"update_time"`
	IsDeleted bool   `json:"is_deleted"`
}