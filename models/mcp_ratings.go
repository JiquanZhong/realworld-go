package models

import "time"

type McpRating struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"uniqueIndex:idx_user_mcp_service" json:"user_id" binding:"required"`
	McpServiceID uint      `gorm:"uniqueIndex:idx_user_mcp_service" json:"mcp_service_id" binding:"required"`
	Stars        uint8     `json:"stars" binding:"required,gte=1,lte=5"`
	Comment      string    `json:"comment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (McpRating) TableName() string {
	return "mcp_ratings"
}
