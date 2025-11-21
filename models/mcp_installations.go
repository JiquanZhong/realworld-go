package models

import "time"

type McpInstallation struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	McpServiceID uint      `gorm:"uniqueIndex:idx_install_user_mcp;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"mcp_service_id" binding:"required"`
	UserID       uint      `gorm:"uniqueIndex:idx_install_user_mcp;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id" binding:"required"`
	LastUsedAt   time.Time `json:"last_used_at"`
	CreatedAt    time.Time `json:"created_at"`
}

func (McpInstallation) TableName() string {
	return "mcp_installations"
}
