package models

import "time"

type Book struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"type:varchar(255)"`
	Description string     `json:"description" gorm:"type:varchar(255)"`
	ImageURL    string     `json:"image_url" gorm:"type:varchar(255)"`
	ReleaseYear int        `json:"release_year"`
	Price       int        `json:"price"`
	TotalPage   int        `json:"total_page"`
	Thickness   string     `json:"thickness" gorm:"type:varchar(10)"` // tipis / tebal

	CategoryID uint      `json:"category_id"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryID"`

	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  string     `json:"created_by" gorm:"type:varchar(100)"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string    `json:"modified_by" gorm:"type:varchar(100)"`
}
