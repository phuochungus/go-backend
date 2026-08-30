package po

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id;type:int;not null;primaryKey; autoIncrement; comment:'Primary Key is Id'"`
	RoleName string `gorm:"column:role_name"`
	IsActive bool   `gorm:"column:is_active"`
	RoleNote string `gorm:"column:role_note"`
}

func (r *Role) TableName() string {
	return "roles"
}
