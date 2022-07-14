package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//User 用户模型
type User struct {
	gorm.Model
	Avatar         string `gorm:"type:varchar(255);"`
	Email          string `gorm:"type:varchar(255);"`
	UserName       string `gorm:"unique;type:varchar(255);"`
	NickName       string `gorm:"type:varchar(255);"`
	Password       string `gorm:"type:varchar(255);"`
	Administrator  bool   //管理员或者普通用户
	PasswordDigest string `gorm:"type:varchar(255);"`
	LastLoginTime  time.Time
	LoginProvince  string  `gorm:"type:varchar(255);"`
	LoginCity      string  `gorm:"type:varchar(255);"`
	LoginLat       float64 // 登录地点维度
	LoginLng       float64 // 登录地点经度
}

const (
	PassWordCost = 12 //密码加密难度
)

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
