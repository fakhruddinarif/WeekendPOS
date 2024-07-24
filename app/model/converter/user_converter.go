package converter

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Photo:     user.Photo,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
func UserToTokenResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		Token: user.Token,
	}
}

func UserToEmployeesResponse(employees *[]entity.User) *model.UserResponse {
	employeesResponse := make([]model.UserResponse, 0)
	for _, e := range *employees {
		employeesResponse = append(employeesResponse, model.UserResponse{
			ID:        e.ID,
			Username:  e.Username,
			Name:      e.Name,
			Email:     e.Email,
			Phone:     e.Phone,
			Photo:     e.Photo,
			Role:      e.Role,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
			DeletedAt: e.DeletedAt,
		})
	}
	return &model.UserResponse{
		Employees: &employeesResponse,
	}
}

func UserToEvent(user *entity.User) *model.UserEvent {
	return &model.UserEvent{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Photo:     user.Photo,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
