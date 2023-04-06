package models

import (
	"fast_gin/models/ctype"
)

// UserModel 用户表
type UserModel struct {
	MODEL
	NickName string     `gorm:"size:36" json:"nick_name"`     // 昵称
	UserName string     `gorm:"size:36" json:"user_name"`     // 用户名
	Password string     `gorm:"size:128" json:"-"`            // 密码
	Avatar   string     `gorm:"size:256" json:"avatar"`       // 头像
	Email    string     `gorm:"size:128" json:"email"`        // 邮箱
	Tel      string     `gorm:"size:18" json:"tel"`           // 手机号
	Addr     string     `gorm:"size:64" json:"addr"`          // 地址
	Token    string     `gorm:"size:64" json:"token"`         // 其他平台的唯一id
	IP       string     `gorm:"size:20" json:"ip"`            // ip地址
	Role     ctype.Role `gorm:"size:4;default:2" json:"role"` // 权限  1 管理员  2 普通用户  3 游客
}
