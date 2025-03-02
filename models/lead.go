package models

type Lead struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(100);not null" json:"name"`
	Email     string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Phone     string `gorm:"type:varchar(20);unique;not null" json:"phone"`
	Address   string `gorm:"not null" json:"address"`
	Status    string `gorm:"type:varchar(50);not null;default:'new';check:status IN ('new', 'approved', 'rejected')" json:"status"`
	CreatedBy uint   `gorm:"not null" json:"created_by"` //FK userID
	CreatedAt string `gorm:"type:timestamptz;default:null" json:"created_at"`
	UpdatedAt string `gorm:"type:timestamptz;default:null" json:"updated_at"`
}
