package models

import "time"

type Project struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	LeadID     uint      `gorm:"not null" json:"lead_id"`                      //FK leadID
	AssignedTo uint      `gorm:"not null" json:"assign_to"`                    //FK userID (sales)
	ProductID  string    `gorm:"type:varchar(100);not null" json:"product_id"` //Array [1,3,4]
	Status     string    `gorm:"type:varchar(50);not null;default:'pending';check:status IN ('pending', 'approved', 'rejected')" json:"status"`
	ApprovalBy uint      `gorm:"default:null" json:"approval_by"` //FK userID (manager)
	ApprovedAt time.Time `gorm:"type:timestamp" json:"approved_at"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:null" json:"updated_at"`
}
