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
func (rq *repoQuery) GetUser(email string) (domain.Core, error) {
	var resQry User
	//var err error
	if err := rq.db.First(&resQry, "email = ? ", email).Error; err != nil {
		log.Error("Error on get user", err.Error())
		return domain.Core{}, err
	}

	// resQry.Token, err = jwt.GenerateJWTToken(resQry.ID)
	// if err != nil {
	// 	return domain.Core{}, err
	// }

	// if err := rq.db.Save(resQry).Error; err != nil {
	// 	return domain.Core{}, err
	// }

	loginUser := ToDomain(resQry)
	return loginUser, nil

}

// Delete implements domain.Repository
func (*repoQuery) Delete(ID uint) (domain.Core, error) {
	panic("unimplemented")
}

// Get implements domain.Repository
func (*repoQuery) Get(email string) (domain.Core, error) {
	panic("unimplemented")
}

// Update implements domain.Repository
func (*repoQuery) Update(updatedData domain.Core, ID uint) (domain.Core, error) {
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

// // Update implements domain.Repository
// func (rq *repoQuery) Update(updatedData domain.Core, ID uint) (domain.Core, error) {
// 	var res User
// 	if err := rq.db.First(&res, "id=?", res.ID).Error; err != nil {
// 		return domain.Core{}, err
// 	}

// 	res.Name = updatedData.Name
// 	res.HP = updatedData.HP
// 	res.Addres = updatedData.Addres

// 	if err := rq.db.Save(&res).Error; err != nil {
// 		return domain.Core{}, err
// 	}
// 	return ToDomain(res), nil
// }

// func (rq *repoQuery) Delete(ID uint) (domain.Core, error) {
// 	var res User
// 	if err := rq.db.First(&res, "id=?", ID).Error; err != nil {
// 		return domain.Core{}, err
// 	}
// 	if err := rq.db.Delete(&res).Error; err != nil {
// 		return domain.Core{}, err
// 	}
// 	return ToDomain(res), nil
// }
