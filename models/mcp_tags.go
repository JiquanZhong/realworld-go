package models

import "time"

type McpTag struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`

	// 多对多关系：一个标签可以被多个 MCP 使用
	McpServices []McpService `gorm:"many2many:mcp_service_tags;foreignKey:ID;joinForeignKey:TagId;References:ID;joinReferences:McpId" json:"mcp_services,omitempty"`
}

func (McpTag) TableName() string {
	return "mcp_tags"
}
