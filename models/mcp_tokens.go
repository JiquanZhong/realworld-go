package models

import "time"

type McpToken struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Token        string    `gorm:"unique;not null" json:"token" binding:"required"`
	UserID       uint      `gorm:"not null" json:"user_id" binding:"required"`
	McpServiceID uint      `gorm:"not null" json:"mcp_service_id" binding:"required"`
	Scope        string    `gorm:"not null" json:"scope" binding:"required"`
	ExpiresAt    time.Time `json:"expires_at" binding:"required"`
	Revoked      bool      `gorm:"default:false" json:"revoked"`
	CreatedAt    time.Time `json:"created_at"`
	LastUsedAt   time.Time `json:"last_used_at"`
}

func (McpToken) TableName() string {
	return "mcp_tokens"
}
