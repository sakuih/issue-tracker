package models

type Issue struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	Description string `json:"title" binding:"min=20,max=500" gorm:"type:varchar(500)"`
	Severity    int8   `gorm:"type:int8"`
	User        User   `json:"reporter" binding:"required" gorm:"foreignkey:UserID"`
}
