package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Email    string
	Password string
	Name     string
	Phone    string
	Bio      string
	Gender   string
	Location string
}

type Repository interface {
	Login(email string) (Core, error)
	Insert(newUser Core) (Core, error)
	Update(updatedData Core) (Core, error)
	Get(ID uint) (Core, error)
	GetAll() ([]Core, error)
	Delete(ID uint) (Core, error)
	// GetUser(newUser Core) (Core, error)
}

type Service interface {
	AddUser(newUser Core) (Core, error)
	UpdateProfile(updatedData Core) (Core, error)
	Profile(ID uint) (Core, error)
	ShowAllUser() ([]Core, error)
	DeleteUser(ID uint) (Core, error)
	Login(email string, password string) (Core, string, error)
}

type Handler interface { // handle reouting req res
	AddUser() echo.HandlerFunc
	ShowAllUser() echo.HandlerFunc
	Login() echo.HandlerFunc
}
