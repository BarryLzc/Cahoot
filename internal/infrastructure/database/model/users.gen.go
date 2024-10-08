// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUsers = "users"

// Users mapped from table <users>
type Users struct {
	ID        int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt *time.Time     `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	Name      *string        `gorm:"column:name;type:longtext" json:"name"`
	Phone     *string        `gorm:"column:phone;type:longtext" json:"phone"`
	Password  *string        `gorm:"column:password;type:longtext" json:"password"`
}

// TableName Users's table name
func (*Users) TableName() string {
	return TableNameUsers
}
