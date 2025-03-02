package models

type Project struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	LeadID     uint   `gorm:"not null" json:"lead_id"`                      //FK leadID
	ProductID  string `gorm:"type:varchar(100);not null" json:"product_id"` //Array [1,3,4]
	Status     string `gorm:"type:varchar(50);not null;default:'pending';check:status IN ('pending', 'approved', 'rejected')" json:"status"`
	ApprovalBy uint   `gorm:"not null" json:"approval_by"` //FK userID (manager)
	ApprovedAt string `gorm:"type:timestamptz;default:null" json:"approved_at"`
	CreatedAt  string `gorm:"type:timestamptz;default:null" json:"created_at"`
	UpdatedAt  string `gorm:"type:timestamptz;default:null" json:"updated_at"`
}
