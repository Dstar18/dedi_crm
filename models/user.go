package models

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(100);not null" json:"name"`
	Email     string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Role      string `gorm:"type:varchar(50);not null;check:role IN ('admin', 'manager', 'sales')" json:"role"`
	CreatedAt string `gorm:"type:timestamptz;default:null" json:"created_at"`
	UpdatedAt string `gorm:"type:timestamptz;default:null" json:"updated_at"`
}
