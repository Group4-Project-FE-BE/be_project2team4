package services

import (
	"be_project2team4/config"
	"be_project2team4/feature/user/domain"
	"be_project2team4/utils/jwt"
	"errors"
	loggo "log"
	"strings"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func New(repo domain.Repository) domain.Service {
	return &repoService{
		qry: repo,
	}
}

type repoService struct {
	qry domain.Repository
}

// Register implements domain.Service
func (rs *repoService) Register(newUser domain.Core) (domain.Core, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("cannot encript password")
	}

	newUser.Password = string(generate)
	res, err := rs.qry.Insert(newUser)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New(config.DUPLICATED_DATA)
		}

		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil

}

// Login implements domain.Service
func (rs *repoService) Login(email, password string) (domain.Core, string, error) {

	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return domain.Core{}, "", errors.New("hp or password empty")
	}

	res, err := rs.qry.GetUser(email)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, "", errors.New("Failed. Error database.")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, "", errors.New("Failed. HP or Password not found.")
		} else {
			return domain.Core{}, "", errors.New("Failed. Process error. Please contact Admin")
		}
	} else {
		loggo.Println("res pass", res.Password, "\n\npass", password)
		err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
		if err != nil {
			return domain.Core{}, "", errors.New("Failed. Incorrect Password.")
		}

		token, err := jwt.GenerateJWTToken(res.ID)
		if err != nil {
			return domain.Core{}, "", err
		}
		return res, token, nil
	}

}

// DeleteProfile implements domain.Service
func (*repoService) DeleteProfile(ID uint) (domain.Core, error) {
	panic("unimplemented")
}

// Profile implements domain.Service
func (*repoService) Profile(email string) (domain.Core, error) {
	panic("unimplemented")
}

// UpdateProfile implements domain.Service
func (*repoService) UpdateProfile(updatedData domain.Core, ID uint) (domain.Core, error) {
	panic("unimplemented")
}

// // AddUser implements domain.Service
// func (rs *repoService) AddUser(newUser domain.Core) (domain.Core, error) {
// 	res, err := rs.qry.Insert(newUser)
// 	if err != nil {
// 		log.Error(err.Error())
// 		if strings.Contains(err.Error(), "duplicate") {
// 			return domain.Core{}, errors.New(config.DUPLICATED_DATA)
// 		}
// 		return domain.Core{}, errors.New("some problem on database")
// 	}
// 	return res, nil
// }

// // ShowAllUser implements domain.Service
// func (rs *repoService) ShowAllUser() ([]domain.Core, error) {
// 	res, err := rs.qry.GetAll()

// 	if err == gorm.ErrRecordNotFound {
// 		log.Error(err.Error())
// 		return nil, gorm.ErrRecordNotFound
// 	} else if err != nil {
// 		log.Error(err.Error())
// 		return nil, errors.New(config.DATABASE_ERROR)
// 	}

// 	if len(res) == 0 {
// 		log.Info("no data")
// 		return nil, errors.New(config.DATA_NOTFOUND)
// 	}
// 	return res, nil
// }

// // Profile implements domain.Service
// func (rs *repoService) Profile(ID uint) (domain.Core, error) {
// 	res, err := rs.qry.Get(ID)
// 	if err != nil {
// 		log.Error(err.Error())
// 		if err == gorm.ErrRecordNotFound {
// 			return domain.Core{}, gorm.ErrRecordNotFound
// 		} else {
// 			return domain.Core{}, errors.New(config.DATABASE_ERROR)
// 		}
// 	}
// 	return res, nil
// }

// // UpdateProfile implements domain.Service
// func (rs *repoService) UpdateProfile(updatedData domain.Core, ID uint) (domain.Core, error) {
// 	res, err := rs.qry.Update(updatedData, ID)
// 	if err != nil {
// 		log.Error(err.Error())
// 		if err == gorm.ErrRecordNotFound {
// 			return domain.Core{}, gorm.ErrRecordNotFound
// 		} else {
// 			return domain.Core{}, errors.New(config.DATABASE_ERROR)
// 		}
// 	}
// 	return res, nil
// }

// // DeleteUser implements domain.Service
// func (rs *repoService) DeleteUser(ID uint) (domain.Core, error) {
// 	res, err := rs.qry.Delete(ID)
// 	if err != nil {
// 		log.Error(err.Error())
// 		if err == gorm.ErrRecordNotFound {
// 			return res, gorm.ErrRecordNotFound
// 		} else {
// 			return res, errors.New(config.DATABASE_ERROR)
// 		}
// 	}
// 	return domain.Core{}, nil
// }
