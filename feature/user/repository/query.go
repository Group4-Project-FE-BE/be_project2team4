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
