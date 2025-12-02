package models

import (
	"time"
)

// TopicStatus 定义主题状态类型
type TopicStatus string

const (
	TopicStatusNormal   TopicStatus = "normal"
	TopicStatusLocked   TopicStatus = "locked"
	TopicStatusArchived TopicStatus = "archived"
)

type ForumTopics struct {
	ID             uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	BoardID        uint        `gorm:"not null" json:"board_id"`
	UserID         uint        `gorm:"not null" json:"user_id"`
	Title          string      `gorm:"type:varchar(200);not null" json:"title"`
	Status         TopicStatus `gorm:"type:varchar(20);default:'normal'" json:"status"`
	ViewCount      int         `gorm:"default:0" json:"view_count"`
	ReplyCount     int         `gorm:"default:0" json:"reply_count"`
	LastPostAt     time.Time   `json:"last_post_at"`
	LastPostUserID uint        `json:"last_post_user_id"`
	IsDeleted      bool        `gorm:"default:false" json:"is_deleted"`

	// 关联关系
	Board ForumBoards `gorm:"foreignKey:BoardID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
	User  User        `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:RESTRICT" json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ForumTopics) TableName() string {
	return "forum_topics"
}
