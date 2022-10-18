package repository

import (
	"be_project2team4/feature/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Phone    string
	Addres   string
	Bio      string
	Gender   string
	Location string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Name:     du.Name,
		Email:    du.Email,
		Phone:    du.Phone,
		Addres:   du.Addres,
		Bio:      du.Bio,
		Gender:   du.Gender,
		Location: du.Location,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Phone:    u.Phone,
		Addres:   u.Addres,
		Bio:      u.Bio,
		Gender:   u.Gender,
		Location: u.Location,
	}
}

func ToDomainArray(au []User) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{ID: val.ID, Name: val.Name, Email: val.Email, Phone: val.Phone, Addres: val.Addres,
			Bio: val.Bio, Gender: val.Gender, Location: val.Location})
	}

	return res
}
