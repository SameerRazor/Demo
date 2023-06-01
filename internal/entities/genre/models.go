package genre

type Genre struct {
	ID        int    `gorm:"primaryKey;autoIncrement Index:idx_genreid(255);not null" json:"id"`
	Genre     string `gorm:"index:idx_genrename(255);not null;type:varchar(100)" json:"genre"`
	CreatedAt string `gorm:"create_time" json:"created_at"`
	UpdatedAt string `gorm:"update_time" json:"updated_at"`
	IsDeleted bool   `json:"is_deleted"`
}
