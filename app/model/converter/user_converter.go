package converter

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/model"
)

func UserToResponse(user *entity.User, employees []entity.User) *model.UserResponse {
	employeeResponse := make([]model.UserResponse, 0)
	for i, employee := range employees {
		employeeResponse[i] = model.UserResponse{
			ID:        employee.ID,
			Username:  employee.Username,
			Name:      employee.Name,
			Email:     employee.Email,
			Phone:     employee.Phone,
			Photo:     employee.Photo,
			Role:      employee.Role,
			CreatedAt: employee.CreatedAt,
			UpdatedAt: employee.UpdatedAt,
			DeletedAt: employee.DeletedAt,
		}
	}

	return &model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Photo:     user.Photo,
		Role:      user.Role,
		Employee:  &employeeResponse,
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
