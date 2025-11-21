package services

import (
	"errors"

	"github.com/JiquanZhong/realworld-go/db"
	"github.com/JiquanZhong/realworld-go/models"
	"github.com/JiquanZhong/realworld-go/utils"
)

type UserService interface {
	CreateUser(req CreateUserRequest) (models.UserResponse, error)
	ListUsers(page, pageSize uint) (Pagination, error)
	GetUser(id uint) (models.UserResponse, error)
	UpdateUser(id int, req UpdateUserRequest) (models.UserResponse, error)
	DeleteUser(id uint) error
	UpdateUserRole(id uint, req UpdateUserRoleRequest) (models.UserResponse, error)
	Login(email, password string) (LoginResponse, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      uint   `json:"age" binding:"gte=0,lte=130"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      uint   `json:"age" binding:"gte=0,lte=130"`
	Role     string `json:"role" binding:"required,oneof=admin user" example:"admin"`
}

type UpdateUserRoleRequest struct {
	Role string `json:"role" binding:"required,oneof=admin user" example:"admin"`
}

type Pagination struct {
	Total    uint                  `json:"total"`
	Page     uint                  `json:"page"`
	PageSize uint                  `json:"page_size"`
	List     []models.UserResponse `json:"list"`
}

type LoginResponse struct {
	Token string              `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  models.UserResponse `json:"user"`
}

func (s *userService) CreateUser(req CreateUserRequest) (models.UserResponse, error) {
	user := models.User{
		Email: req.Email,
		Name:  req.Name,
		Age:   req.Age,
		Role:  models.RoleUser,
	}
	user.SetPassword(req.Password)

	if err := db.GetDB().Create(&user).Error; err != nil {
		return models.UserResponse{}, err
	}
	return user.ToResponse(), nil
}

func (s *userService) ListUsers(page, pageSize uint) (Pagination, error) {
	var users []models.User

	offset := (page - 1) * pageSize

	var total int64
	db.GetDB().Model(&models.User{}).Count(&total)
	db.GetDB().Offset(int(offset)).Limit(int(pageSize)).Find(&users)

	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user.ToResponse())
	}
	return Pagination{
		Total:    uint(total),
		Page:     page,
		PageSize: pageSize,
		List:     userResponses,
	}, nil
}

func (s *userService) GetUser(id uint) (models.UserResponse, error) {
	var user models.User
	if err := db.GetDB().First(&user, id).Error; err != nil {
		return models.UserResponse{}, err
	}
	return user.ToResponse(), nil
}

func (s *userService) UpdateUser(id int, req UpdateUserRequest) (models.UserResponse, error) {
	var user models.User

	if err := db.GetDB().First(&user, id).Error; err != nil {
		return models.UserResponse{}, err
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Age = req.Age
	user.SetPassword(req.Password)

	if err := db.GetDB().Save(&user).Error; err != nil {
		return models.UserResponse{}, err
	}
	return user.ToResponse(), nil
}

func (s *userService) DeleteUser(id uint) error {
	if err := db.GetDB().Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *userService) UpdateUserRole(id uint, req UpdateUserRoleRequest) (models.UserResponse, error) {
	user := &models.User{
		ID:   id,
		Role: req.Role,
	}
	if err := db.GetDB().Save(&user).Error; err != nil {
		return models.UserResponse{}, err
	}
	return user.ToResponse(), nil
}

func (s *userService) Login(email, password string) (LoginResponse, error) {
	var user models.User

	if err := db.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
		return LoginResponse{}, err
	}

	if !user.CheckPassword(password) {
		return LoginResponse{}, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, user.Name, user.Email, user.Role)
	if err != nil {
		return LoginResponse{}, err
	}
	return LoginResponse{Token: token, User: user.ToResponse()}, nil
}
