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
type EditUserRequestFormat struct {
	Name     string `json:"name"  form:"name"`
	Email    string `json:"email"  form:"email"`
	Password string `json:"password"  form:"password"`
	Phone    string `json:"phone"  form:"phone"`
	Bio      string `json:"bio"  form:"bio"`
	Gender   string `json:"gender"  form:"gender"`
	Location string `json:"location"  form:"location"`
}
type GetUserRequestFormat struct {
	Email string `json:"email"  form:"email"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Name: cnv.Name, Email: cnv.Email, Password: cnv.Password}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Email: cnv.Email, Password: cnv.Password}
	case EditUserRequestFormat:
		cnv := i.(EditUserRequestFormat)
		return domain.Core{
			Name:     cnv.Name,
			Email:    cnv.Email,
			Password: cnv.Password,
			Phone:    cnv.Phone,
			Bio:      cnv.Bio,
			Gender:   cnv.Gender,
			Location: cnv.Location,
		}
	case GetUserRequestFormat:
		cnv := i.(GetUserRequestFormat)
		return domain.Core{Email: cnv.Email}
	}
	return domain.Core{}
}
