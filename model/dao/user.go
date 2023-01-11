package dao

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableNameUsers = "users"
)

// gorm.Model 的定义
type Model struct {
	ID        uint64         `gorm:"column:id;not null;primaryKey"`
	CreatedAt time.Time      `gorm:"column:create_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:update_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;not null;default:''"`
	PassWord string `gorm:"column:password;not null;default:''"`
	Phone    string `gorm:"column:phone;not null;default:'';unique"`
}

func (User) TableName() string {
	return TableNameUsers
}

// 初始化表结构
func NewUserProvider(db *gorm.DB) *UserProvider {
	if err := db.AutoMigrate(new(User)); err != nil {
		panic(err)
	}
	return &UserProvider{db}
}

type UserProvider struct {
	db *gorm.DB
}
