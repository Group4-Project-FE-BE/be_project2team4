package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint
	Name_User string
	Image_Url string
	Content   string
	IDUser    uint
}

type RepositoryInterface interface {
	GetAll() ([]Core, error)
	Get(ID string) (Core, error)
	Insert(newData Core) (Core, error)
	Update(updatedData Core, ID uint) (Core, error)
	Delete(ID string) (Core, error)
}
type ServiceInterface interface {
	GetAll() ([]Core, error)
	Get(ID string) (Core, error)
	Insert(newData Core) (Core, error)
	Update(updatedData Core, ID string) (Core, error)
	Delete(ID string) (Core, error)
	IsAuthorized(c echo.Context) error
}
