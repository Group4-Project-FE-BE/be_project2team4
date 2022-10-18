package repository

import (
	"be_project2team4/feature/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	Name     string
	Phone    string
	Bio      string
	Gender   string
	Location string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Email:    du.Email,
		Password: du.Password,
		Name:     du.Name,
		Phone:    du.Phone,
		Bio:      du.Bio,
		Gender:   du.Gender,
		Location: du.Location,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
		Phone:    u.Phone,
		Bio:      u.Bio,
		Gender:   u.Gender,
		Location: u.Location,
	}
}

func ToDomainArray(au []User) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{ID: val.ID, Name: val.Name, HP: val.HP, Addres: val.Addres})
	}

	return res
}
