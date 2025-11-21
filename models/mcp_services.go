package models

import (
	"time"

	"gorm.io/gorm"
)

type McpService struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"not null" json:"name" binding:"required"`
	IconUrl      string         `json:"icon_url"`
	Description  string         `json:"description"`
	Endpoint     string         `json:"endpoint" binding:"required,url"`
	JsonSchema   string         `gorm:"type:text" json:"json_schema"`
	Category     string         `json:"category"`
	InstallCount uint           `json:"install_count"`
	RatingAvg    float32        `json:"rating_avg"`
	IsVisible    bool           `gorm:"default:true" json:"is_visible"`
	SubmitterID  uint           `json:"submitter_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeleteAt     gorm.DeletedAt `gorm:"index" json:"delete_at"`

	// 多对多关系：一个 MCP 可以有多个标签
	Tags []McpTag `gorm:"many2many:mcp_service_tags;foreignKey:ID;joinForeignKey:McpId;References:ID;joinReferences:TagId" json:"tags,omitempty"`
}

func (McpService) TableName() string {
	return "mcp_services"
}