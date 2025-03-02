package models

type Product struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Desciprtion string `gorm:"not null" json:"description"`
	Price       string `gorm:"type:varchar(100);not null;default:'0'" json:"price"`
}
