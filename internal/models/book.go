package models

type Book struct {
	ID              int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title           string    `gorm:"not null" json:"title"`
	AuthorID        int       `gorm:"not null" json:"author_id"`
	PublicationDate int       `gorm:"not null" json:"date"`
	GenreID         int       `gorm:"not null" json:"genre_id"`
}
