package models

// PriorityStatus represents the structure of the priority_status table
type PriorityStatus struct {
	PriorityId     int     `gorm:"column:priority_id" json:"priority_id"`
	PriorityStatus *string `gorm:"column:priority_status" json:"priority_status"`
}

func (PriorityStatus) TableName() string {
	return "priority_status"
}
