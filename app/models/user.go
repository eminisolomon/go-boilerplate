package models

type User struct {
	Cod        int64  `json:"id"    gorm:"column:usr_cod; primaryKey; <-:create"`
	ExternalID string `json:"external_id" gorm:"column:external_id; unique; not null"`
	Firstname  string `json:"firstname" gorm:"column:firstname; type:text"`
	Lastname   string `json:"lastname" gorm:"column:lastname; type:varchar(255)"`
	Username   string `json:"username" gorm:"column:username; type:varchar(255); not null"`
	Email      string `json:"email" gorm:"column:email; uniqueIndex; not null"`
	Password   string `json:"password" gorm:"column:password; not null"`
	Phone      string `json:"phone" gorm:"column:phone; type:text"`
	Role       string `json:"role" gorm:"column:role; not null; size:255; default:'user'"`
}
