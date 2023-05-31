package book

type Book struct {
	ID              int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title           string `gorm:"uniqueIndex:idx_title_author_genre_pubdate(255);not null;type:varchar(100)" json:"title"`
	AuthorName      string `gorm:"uniqueIndex:idx_title_author_genre_pubdate(255);not null;type:varchar(100)" json:"author_name"`
	AuthorId        int    `gorm:"not null" json:"author_id"`
	GenreName       string `gorm:"uniqueIndex:idx_title_author_genre_pubdate(255);not null;type:varchar(100)" json:"genre_name"`
	GenreId         int    `gorm:"not null" json:"genre_id"`
	PublicationDate string `gorm:"uniqueIndex:idx_title_author_genre_pubdate(255);not null;type:varchar(100)" json:"publication_date"`
	CreatedAt       string `gorm:"create_time" json:"created_at"`
	UpdatedAt       string `gorm:"update_time" json:"updated_at"`
	IsDeleted       bool   `json:"is_deleted"`
}
