package converter

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/model"
)

func EmployeeToResponse(employee *entity.Employee) *model.EmployeeResponse {
	return &model.EmployeeResponse{
		ID:        employee.ID,
		Name:      employee.Name,
		Email:     employee.Email,
		Username:  employee.Username,
		Phone:     employee.Phone,
		Address:   employee.Address,
		Photo:     employee.Photo,
		Token:     employee.Token,
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
	}
}

func EmployeeToEvent(employee *entity.Employee) *model.EmployeeEvent {
	return &model.EmployeeEvent{
		ID:        employee.ID,
		Name:      employee.Name,
		Email:     employee.Email,
		Username:  employee.Username,
		Phone:     employee.Phone,
		Address:   employee.Address,
		Photo:     employee.Photo,
		Token:     employee.Token,
		UserID:    employee.UserId,
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
	}
}
