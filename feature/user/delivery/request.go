package delivery

import "be_project2team4/feature/user/domain"

type RegisterFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
}
type editUserFormat struct {
	Name   string `json:"name"  form:"name"`
	HP     string `json:"hp" form:"hp"`
	Addres string `json:"addres"  form:"addres"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Name: cnv.Name, Email: cnv.Email, Password: cnv.Password}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Email: cnv.Email, Password: cnv.Password}

	}
	return domain.Core{}
}
