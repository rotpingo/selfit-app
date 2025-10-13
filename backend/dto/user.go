package dto

import (
	"selfit/models"
	"time"
)

type UserResponseDTO struct {
	ID        int64     `json:"id" binding:"required"`
	Name      string    `json:"name"`
	Email     string    `json:"email" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserAuthDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserDTO struct {
	ID    int64  `json:"id" binding:"required"`
	Name  string `json:"name"`
	Email string `json:"email" binding:"required"`
}

// func (dto UpdateUserDTO) ToUserModel(userId int64) *models.User {
// 	return &models.User{
// 		ID:        dto.ID,
// 		Name:      dto.Name,
// 		Email:     dto.Email,
// 		UpdatedAt: time.Now(),
// 	}
// }
//
// func UserToResponseDTO(user *models.User) UserResponseDTO {
// 	return UserResponseDTO{
// 		ID:    user.ID,
// 		Name:  user.Name,
// 		Email: user.Email,
// 	}
// }

func (dto UserAuthDTO) ToUserModel() *models.User {
	return &models.User{
		Email:    dto.Email,
		Password: dto.Password,
	}
}
