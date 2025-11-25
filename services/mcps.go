package services

import (
	"github.com/JiquanZhong/realworld-go/db"
	"github.com/JiquanZhong/realworld-go/models"
	"github.com/JiquanZhong/realworld-go/utils"
	"gorm.io/gorm"
)

type McpService interface {
	ListMcpServices(page, pageSize uint, listOptions utils.ListOptions, searchKeyword string) (utils.Pagination, error)
	RegisterMcpService(req RegisterMcpServiceRequest) (models.McpService, error)
	GetMcpService(id uint) (models.McpService, error)
	DeleteMcpService(id uint) error
}

type mcpService struct{}

func NewMcpService() McpService {
	return &mcpService{}
}

type RegisterMcpServiceRequest struct {
	Name        string `json:"name" binding:"required"`
	IconUrl     string `json:"icon_url"`
	Description string `json:"description" binding:"required"`
	Endpoint    string `json:"endpoint" binding:"required,url"`
	JsonSchema  string `json:"json_schema"`
	Category    string `json:"category"`
	Tags        []uint `json:"tags"`
	SubmitterID uint   `json:"submitter_id"`
}

func (s *mcpService) ListMcpServices(page, pageSize uint, listOptions utils.ListOptions, searchKeyword string) (utils.Pagination, error) {

	var mcpServices []models.McpService

	offset := (page - 1) * pageSize

	dbQuery := db.GetDB().Model(&models.McpService{})

	// 添加搜索关键词过滤
	if searchKeyword != "" {
		searchPattern := "%" + searchKeyword + "%"
		dbQuery = dbQuery.Where("name LIKE ? OR description LIKE ? OR category LIKE ?",
			searchPattern, searchPattern, searchPattern)
	}

	if listOptions.By != "" {
		order := listOptions.By
		if !listOptions.Asc {
			order += " desc"
		}
		dbQuery = dbQuery.Order(order)
	}

	var total int64
	// 注意：计算总数时也要应用搜索条件
	countQuery := db.GetDB().Model(&models.McpService{})
	if searchKeyword != "" {
		searchPattern := "%" + searchKeyword + "%"
		countQuery = countQuery.Where("name LIKE ? OR description LIKE ? OR category LIKE ?",
			searchPattern, searchPattern, searchPattern)
	}
	countQuery.Count(&total)

	err := dbQuery.Preload("Tags").Offset(int(offset)).Limit(int(pageSize)).Find(&mcpServices).Error
	if err != nil {
		return utils.Pagination{}, err
	}

	var mcpResponses []models.McpResponse
	for _, mcpService := range mcpServices {
		mcpResponses = append(mcpResponses, mcpService.ToResponse())
	}

	return utils.Pagination{
		Total:    uint(total),
		Page:     page,
		PageSize: pageSize,
		List:     &mcpResponses,
	}, nil

}

func (s *mcpService) RegisterMcpService(req RegisterMcpServiceRequest) (models.McpService, error) {
	mcp := models.McpService{
		Name:        req.Name,
		IconUrl:     req.IconUrl,
		Description: req.Description,
		Endpoint:    req.Endpoint,
		JsonSchema:  req.JsonSchema,
		Category:    req.Category,
		SubmitterID: req.SubmitterID,
	}

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&mcp).Error; err != nil {
			return err
		}

		if len(req.Tags) > 0 {
			var tags []models.McpTag
			if err := tx.Where("id IN ?", req.Tags).Find(tags).Error; err != nil {
				return err
			}

			if err := tx.Model(&mcp).Association("Tags").Replace(tags); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return models.McpService{}, err
	}

	return s.GetMcpService(mcp.ID)
}

func (s *mcpService) GetMcpService(id uint) (models.McpService, error) {
	var mcp models.McpService

	if err := db.GetDB().Preload("Tags").First(&mcp, id).Error; err != nil {
		return models.McpService{}, err
	}

	return mcp, nil
}

func (s *mcpService) DeleteMcpService(id uint) error {

	if err := db.GetDB().Delete(&models.McpService{}, id).Error; err != nil {
		return err
	}
	return nil
}
