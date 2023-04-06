package models

import "time"

type MODEL struct {
	ID        uint      `json:"id" gorm:"primarykey"` // 主键ID
	CreatedAt time.Time `json:"date"`                 // 创建时间
	UpdatedAt time.Time `json:"-"`                    // 更新时间
}
