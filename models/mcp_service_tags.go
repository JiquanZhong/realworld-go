package models

// MCPServiceTag 表示 MCP 服务和标签之间的多对多关系
type MCPServiceTag struct {
	McpId uint `gorm:"primaryKey;column:mcp_id" json:"mcp_id"`
	TagId uint `gorm:"primaryKey;column:tag_id" json:"tag_id"`

	// 外键关联
	McpService McpService `gorm:"foreignKey:McpId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Tag        McpTag     `gorm:"foreignKey:TagId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

func (MCPServiceTag) TableName() string {
	return "mcp_service_tags"
}