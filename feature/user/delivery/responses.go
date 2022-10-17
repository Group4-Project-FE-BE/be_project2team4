package delivery

import "be_project2team4/feature/user/domain"

func SuccessResponses(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponses(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type registerRespons struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type editUserRespons struct {
	Name   string `json:"name"`
	HP     string `json:"hp"`
	Addres string `json:"addres"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = registerRespons{ID: cnv.ID, Name: cnv.Name, Email: cnv.Email}
	case "edit":
		cnv := core.(domain.Core)
		res = editUserRespons{Name: cnv.Name}
	}
	return res
}
