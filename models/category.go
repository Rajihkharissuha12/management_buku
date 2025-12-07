package models

import "time"

type Category struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `gorm:"type:varchar(255)"`
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt *time.Time
	ModifiedBy *string

	Books []Book `gorm:"foreignKey:CategoryID"`
}
