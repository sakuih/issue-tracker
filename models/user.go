package models

type User struct {
	ID       uint64  `gorm:"primary_key;auto-increment" json:"id"`
	email    string  `json:"email" binding:"required,email" gorm:"type:varchar(100)"`
	password string  `json:"email" binding:"required,email" gorm:"type:varchar(100)`
	issues   []Issue `json:"issues"`
}
