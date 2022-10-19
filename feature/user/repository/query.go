package repository

import (
	"be_project2team4/feature/user/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

// Insert implements domain.Repository
func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {
	var cnv User

	cnv = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("Error on insert user", err.Error())
		return domain.Core{}, err
	}

	// selesai dari DB
	return ToDomain(cnv), nil
}

// GetUser implements domain.Repository
// get all user data to show user
func (rq *repoQuery) GetUser(email string) (domain.Core, error) {
	var resQry User
	//var err error
	if err := rq.db.First(&resQry, "email = ? ", email).Error; err != nil {
		log.Error("Error on get user", err.Error())
		return domain.Core{}, err
	}

	loginUser := ToDomain(resQry)
	return loginUser, nil

}

// Get implements domain.Repository
func (*repoQuery) Get(email string) (domain.Core, error) {
	panic("unimplemented")
}

// // Insert implements domain.Repository
// func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {
// 	var cnv User
// 	cnv = FromDomain(newUser)
// 	if err := rq.db.Create(&cnv).Error; err != nil {
// 		log.Error("Error on insert user", err.Error())
// 		return domain.Core{}, err
// 	}
// 	return ToDomain(cnv), nil
// }

// // GetAll implements domain.Repository
// func (rq *repoQuery) GetAll() ([]domain.Core, error) {
// 	var res []User
// 	if err := rq.db.Find(&res).Error; err != nil {
// 		log.Error("error on get all user", err.Error())
// 		return []domain.Core{}, err
// 	}
// 	return ToDomainArray(res), nil
// }

// // Get implements domain.Repository
// func (rq *repoQuery) Get(ID uint) (domain.Core, error) {
// 	var res User
// 	if err := rq.db.First(&res, "id =?", ID).Error; err != nil {
// 		log.Error("error on getuserid", err.Error())
// 		return domain.Core{}, err
// 	}
// 	return ToDomain(res), nil
// }

// Update implements domain.Repository
func (rq *repoQuery) Update(updatedData domain.Core, ID uint) (domain.Core, error) {
	var template User
	var cnv User
	cnv = FromDomain(updatedData)

	err := rq.db.Where("id = ?", ID).First(&template).Error
	if err != nil {
		return domain.Core{}, err
	}

	cnv.ID = ID
	if err := rq.db.Save(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	return ToDomain(cnv), nil
}

// Delete implements domain.Repository
func (rq *repoQuery) Delete(ID uint) (domain.Core, error) {
	var res User
	if err := rq.db.First(&res, "id=?", ID).Error; err != nil {
		return domain.Core{}, err
	}
	if err := rq.db.Delete(&res).Error; err != nil {
		return domain.Core{}, err
	}
	return ToDomain(res), nil
}
