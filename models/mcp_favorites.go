package models

import "time"

type McpFavorite struct {
	UserID       uint      `gorm:"primaryKey" json:"user_id" binding:"required"`
	McpServiceID uint      `gorm:"primaryKey" json:"mcp_service_id" binding:"required"`
	CreatedAt    time.Time `json:"created_at"`
}

func (McpFavorite) TableName() string {
	return "mcp_favorites"
}
