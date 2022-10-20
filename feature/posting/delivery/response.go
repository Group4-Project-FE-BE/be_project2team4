package delivery

import (
	"be_project2team4/feature/posting/domain"
	"log"
)

type PostingResponseFormat struct {
	ID        uint   `json:"id"`
	Name_User string `json:"name_user"`
	Image_Url string `json:"image_url"`
	Content   string `json:"content"`
	IDUser    uint   `json:"id_user"`
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessDeleteResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "posting":
		var posting PostingResponseFormat
		cnv := core.(domain.Core)
		posting = PostingResponseFormat{
			ID:        cnv.ID,
			Name_User: cnv.Name_User,
			Image_Url: cnv.Image_Url,
			Content:   cnv.Content,
			IDUser:    cnv.IDUser,
		}
		res = posting
	case "postings":
		var arrPosting []PostingResponseFormat
		cnv := core.([]domain.Core)
		log.Println("\n\n isi res =", cnv, "\n\n")
		for _, val := range cnv {
			arrPosting = append(arrPosting,
				PostingResponseFormat{
					ID:        val.ID,
					Name_User: val.Name_User,
					Image_Url: val.Image_Url,
					Content:   val.Content,
					IDUser:    val.IDUser,
				})
		}
		res = arrPosting
	case "postingAllComments":
		// pending, nunggu comment dibuat

		// var arrPosting []PostingResponseFormat
		// cnv := core.([]domain.Core)
		// log.Println("\n\n isi res =", cnv, "\n\n")
		// for _, val := range cnv {
		// 	arrPosting = append(arrPosting,
		// 		PostingResponseFormat{
		// 			ID:        val.ID,
		// 			Name_User: val.Name_User,
		// 			Image_Url: val.Image_Url,
		// 			Content:   val.Content,
		// 			IDUser:    val.IDUser,
		// 		})
		// }
		// res = arrPosting
	}

	return res
}
