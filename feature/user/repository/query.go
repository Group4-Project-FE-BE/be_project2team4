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
	return ToDomain(cnv), nil
}

// GetAll implements domain.Repository
// get all user data to show user
func (rq *repoQuery) GetAll() ([]domain.Core, error) {
	var res []User
	if err := rq.db.Find(&res).Error; err != nil {
		log.Error("error on get all user", err.Error())
		return []domain.Core{}, err
	}
	return ToDomainArray(res), nil
}

// Get implements domain.Repository
func (rq *repoQuery) Get(Email string) (domain.Core, error) {
	var res User
	if err := rq.db.First(&res, "email =?", Email).Error; err != nil {
		log.Error("error on getuseremail", err.Error())
		return domain.Core{}, err
	}
	return ToDomain(res), nil
}

// Update implements domain.Repository
func (rq *repoQuery) Update(updatedData domain.Core, ID uint) (domain.Core, error) {
	var res User
	if err := rq.db.First(&res, "id=?", res.ID).Error; err != nil {
		return domain.Core{}, err
	}

	res.Name = updatedData.Name
	res.Email = updatedData.Email
	res.Phone = updatedData.Phone
	res.Addres = updatedData.Addres
	res.Bio = updatedData.Bio
	res.Gender = updatedData.Gender
	res.Location = updatedData.Location

	if err := rq.db.Save(&res).Error; err != nil {
		return domain.Core{}, err
	}
	return ToDomain(res), nil
}

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
