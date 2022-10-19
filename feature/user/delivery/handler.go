package delivery

import (
	"be_project2team4/feature/user/domain"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}

	e.POST("/register", handler.Register())
	e.POST("/login", handler.Login())

}

// registrasi add user
func (us *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.Register(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponses(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponses("berhasil register", ToResponse(res, "reg", "")))
	}

}

func (us *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses("cannot bind input"))
		}
		cnv := ToDomain(input)
		log.Println("\n\n\ndata login \n", input, "\n\n")
		res, token, err := us.srv.Login(cnv.Email, cnv.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponses(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessLoginResponses("berhasil login", ToResponse(res, "login", token)))
	}
}
