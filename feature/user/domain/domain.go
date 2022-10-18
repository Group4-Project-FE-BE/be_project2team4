package domain

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID       uint
	Name     string
	Email    string
	Phone    string
	Addres   string
	Bio      string
	Gender   string
	Location string
}

type Repository interface {
	Insert(newUser Core) (Core, error)              //register
	GetUser(Email string) (Core, error)             //login
	Update(updatedData Core, ID uint) (Core, error) //update/edit
	Get(Email string) (Core, error)                 //get/find
	ShowAllUser() ([]Core, error)                   //show all user
	Delete(ID uint) (Core, error)                   //delete
}

type Service interface {
	Register(newUser Core) (Core, error)
	Login(Email string) (Core, error)
	UpdateProfile(updatedData Core, ID uint) (Core, error)
	ShowAllUser() ([]Core, error)
	Profile(Email string) (Core, error)
	DeleteProfile(ID uint) (Core, error)
}

type Handler interface {
	ShowAllUser() echo.HandlerFunc
	Get() echo.HandlerFunc
}
