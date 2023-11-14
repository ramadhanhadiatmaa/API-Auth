package models

type Auth struct {
	Username string `gorm:"type:varchar(50); primaryKey" json:"username"`
	Password string `gorm:"type:varchar(250)" json:"password"`
	Tipe     string `gorm:"type:varchar(50)" json:"tipe"`
}