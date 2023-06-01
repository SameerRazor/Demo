package author

type Author struct {
	ID          int    `json:"id" gorm:"primaryKey Index:idx_authorid(255);not null;type:varchar(100)"`
	AuthorName  string `json:"author_name" gorm:"uniqueIndex:idx_author_bio_dob_nation(255);not null;type:varchar(100) Index:idx_authorname(255);not null;type:varchar(100)"`
	Biography   string `json:"biography" gorm:"uniqueIndex:idx_author_bio_dob_nation(255);not null;type:varchar(100)"`
	DateOfBirth string `json:"date_of_birth" gorm:"uniqueIndex:idx_author_bio_dob_nation(255);not null;type:varchar(100)"`
	Nationality string `json:"nationality" gorm:"uniqueIndex:idx_author_bio_dob_nation(255);not null;type:varchar(100) Index:idx_nationality(255)"`
	CreatedAt   string `json:"created_at" gorm:"create_time"`
	UpdatedAt   string `json:"updated_at" gorm:"update_time"`
	IsDeleted   bool   `json:"is_deleted"`
}
