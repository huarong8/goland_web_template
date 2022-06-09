package repository

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID    int64     `gorm:"column:user_id;not null" json:"user_id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Email     string    `gorm:"column:email;not null" json:"email"`
	Passwd    string    `gorm:"column:passwd;not null" json:"passwd"`
	Stat      int32     `gorm:"column:stat;not null;default:1" json:"stat"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt  time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP" json:"update_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

func (u *User) FindById() string {
	return "findById funcion....."
}
