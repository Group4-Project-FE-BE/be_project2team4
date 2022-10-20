package delivery

import (
	"be_project2team4/config"
	"be_project2team4/feature/posting/domain"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var key string

type postingHandler struct {
	srv domain.ServiceInterface
}

func InitJWT(c *config.AppConfig) {
	key = c.JWSecret
}

func New(e *echo.Echo, srv domain.ServiceInterface) {
	handler := postingHandler{srv: srv}
	e.POST("/postings", handler.AddPosting(), middleware.JWT([]byte(key)))
	e.GET("/postings", handler.GetAllPosting())
	e.GET("/postings/:id", handler.GetPosting())
	e.GET("/postings/:id/comments", handler.GetPostingAllComment())
	e.PUT("/postings/:id", handler.UpdatePosting(), middleware.JWT([]byte(key)))
	e.DELETE("/postings/:id", handler.DeletePosting(), middleware.JWT([]byte(key)))
}

func (us *postingHandler) GetAllPosting() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := us.srv.GetAll()
		//log.Println("\n\n\n res GET ALL =", res, "\n\n")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success get postings.", ToResponse(res, "postings")))
	}
}

func (us *postingHandler) GetPosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token
		// err := us.srv.IsAuthorized(c)
		// if err != nil {
		// 	return c.JSON(http.StatusUnauthorized, FailResponse(err.Error()))
		// } else {
		// 	log.Println("Authorized request.")
		// }

		paramID := c.Param("id")
		res, err := us.srv.Get(paramID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success get posting.", ToResponse(res, "posting")))
	}
}

func (us *postingHandler) GetPostingAllComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token

		paramID := c.Param("id")
		res, err := us.srv.Get(paramID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success get posting.", ToResponse(res, "posting")))
	}
}

func (us *postingHandler) AddPosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponse(err.Error()))
		} else {
			log.Println("Authorized request.")
		}

		var input PostingInsertRequestFormat
		if err := c.Bind(&input); err != nil {
			log.Println("Error Bind = ", err.Error())
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		log.Println("\n\n\n input posting handler : ", input, "\n\n\n")
		cnv := ToDomain(input)
		res, err := us.srv.Insert(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success add posting.", ToResponse(res, "posting")))
	}

}

func (us *postingHandler) UpdatePosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponse(err.Error()))
		} else {
			log.Println("Authorized request.")
		}

		var input PostingRequestFormat
		paramID := c.Param("id")
		if err := c.Bind(&input); err != nil {
			log.Println("Error Bind = ", err.Error())
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		log.Println("\n\n\nid handler : ", paramID)
		log.Println("\n\n\n input handler : ", input)

		//log.Printf("\n\n\n isi input", &input, "\n\n\n")
		cnv := ToDomain(input)
		//log.Printf("\n\n\n isi cnv", cnv, "\n\n\n")
		res, err := us.srv.Update(cnv, paramID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success update posting.", ToResponse(res, "posting")))
	}
}

func (us *postingHandler) DeletePosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponse(err.Error()))
		} else {
			log.Println("Authorized request.")
		}

		paramID := strings.TrimSpace(c.Param("id"))
		// validasi not empty id
		if paramID == "" {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		} else {
			res, err := us.srv.Get(paramID)
			log.Println("res data :", res)
			// validasi get data by id
			if err != nil {
				log.Println("Error get data. Error :", err.Error())
				return c.JSON(http.StatusInternalServerError, FailResponse("Data not found"))
			} else {
				res2, err2 := us.srv.Delete(paramID)
				log.Println("res2 data :", res2)
				if err2 != nil {
					return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
				} else {
					return c.JSON(http.StatusCreated, SuccessDeleteResponse("Success delete posting."))
				}
			}
		}

	}
}

// func (us *userHandler) UpdateProfile() (domain.Core, error) {

// }
// func (us *userHandler) Profile() (domain.Core, error) {
// 	res, err := us.qry.Get(ID)
// 	if err != nil {
// 		log.Error(err.Error())
// 		if strings.Contains(err.Error(), "table") {
// 			return domain.Core{}, errors.New("database error")
// 		} else if strings.Contains(err.Error(), "found") {
// 			return domain.Core{}, errors.New("no data")
// 		}
// 	}

//		return res, nil
//	}
