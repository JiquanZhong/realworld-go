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
	JsonSchema   string         `gorm:"type:text" json:"-"`
	Category     string         `json:"category"`
	InstallCount uint           `json:"install_count"`
	RatingAvg    float32        `json:"rating_avg"`
	Detail       string         `gorm:"type:text" json:"detail,omitempty"`
	IsVisible    bool           `gorm:"default:true" json:"is_visible"`
	SubmitterID  uint           `json:"submitter_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeleteAt     gorm.DeletedAt `gorm:"index" json:"delete_at"`

	// 多对多关系：一个 MCP 可以有多个标签
	Tags []McpTag `gorm:"many2many:mcp_service_tags;foreignKey:ID;joinForeignKey:McpId;References:ID;joinReferences:TagId" json:"tags,omitempty"`
}

type McpDetailResponse struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	IconUrl       string    `json:"icon_url"`
	Description   string    `json:"description"`
	Endpoint      string    `json:"endpoint"`
	Category      string    `json:"category"`
	InstallCount  uint      `json:"install_count"`
	RatingAvg     float32   `json:"rating_avg"`
	Detail        string    `json:"detail"`
	FavoriteCount uint      `json:"favorite_count"`
	SubmitterID   uint      `json:"submitter_id"`
	SubmitterName string    `json:"submitter_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	// 多对多关系：一个 MCP 可以有多个标签
	Tags []McpTagResponse `json:"tags"`
}

type McpResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	IconUrl      string    `json:"icon_url"`
	Description  string    `json:"description"`
	Endpoint     string    `json:"endpoint"`
	Category     string    `json:"category"`
	InstallCount uint      `json:"install_count"`
	RatingAvg    float32   `json:"rating_avg"`
	SubmitterID  uint      `json:"submitter_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// 多对多关系：一个 MCP 可以有多个标签
	Tags []McpTagResponse `json:"tags"`
}

func (McpService) TableName() string {
	return "mcp_services"
}

func (mcpService *McpService) ToResponse() McpResponse {

	var tags []McpTagResponse
	for _, tag := range mcpService.Tags {
		tags = append(tags, tag.ToResponse())
	}

	return McpResponse{
		ID:           mcpService.ID,
		Name:         mcpService.Name,
		IconUrl:      mcpService.IconUrl,
		Description:  mcpService.Description,
		Category:     mcpService.Category,
		InstallCount: mcpService.InstallCount,
		RatingAvg:    mcpService.RatingAvg,
		CreatedAt:    mcpService.CreatedAt,
		UpdatedAt:    mcpService.UpdatedAt,
		Tags:         tags,
	}
}

func (mcpService *McpService) ToDetailResponseWithDetail() McpDetailResponse {

	var tags []McpTagResponse
	for _, tag := range mcpService.Tags {
		tags = append(tags, tag.ToResponse())
	}

	return McpDetailResponse{
		ID:           mcpService.ID,
		Name:         mcpService.Name,
		IconUrl:      mcpService.IconUrl,
		Description:  mcpService.Description,
		Category:     mcpService.Category,
		InstallCount: mcpService.InstallCount,
		RatingAvg:    mcpService.RatingAvg,
		Detail:       mcpService.Detail,
		SubmitterID:  mcpService.SubmitterID,
		CreatedAt:    mcpService.CreatedAt,
		UpdatedAt:    mcpService.UpdatedAt,
		Tags:         tags,
	}
}
