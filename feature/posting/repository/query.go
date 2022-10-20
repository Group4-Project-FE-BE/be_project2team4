package repository

import (
	"be_project2team4/feature/posting/domain"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.RepositoryInterface {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) Get(ID string) (domain.Core, error) {
	var resQry Posting
	if err := rq.db.Where("ID = ?", ID).First(&resQry).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) GetAll() ([]domain.Core, error) {
	var resQry []Posting
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) Insert(newData domain.Core) (domain.Core, error) {
	var newInput Posting
	newInput = FromDomain(newData)
	if err := rq.db.Create(&newInput).Error; err != nil {
		return domain.Core{}, err
	}

	// convert ke core lg
	newData = ToDomain(newInput)
	return newData, nil
}

func (rq *repoQuery) Update(updatedData domain.Core, ID uint) (domain.Core, error) {
	var currentData Posting

	// validasi jika data tidak ditemukan
	err := rq.db.Where("id = ?", ID).First(&currentData).Error
	if err != nil {
		return domain.Core{}, err
	}

	currentData.ID = ID
	currentData.Name_User = updatedData.Name_User
	currentData.Image_Url = updatedData.Image_Url
	currentData.Content = updatedData.Content
	currentData.IDUser = updatedData.IDUser

	// log.Println("\n\n\n query isi update", updatedBookInput, "\n\n\n")
	if err := rq.db.Where(&currentData.ID).Updates(&currentData.UpdatedAt).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	updatedData = ToDomain(currentData)
	return updatedData, nil
}

func (rq *repoQuery) Delete(ID string) (domain.Core, error) {
	var res Posting
	if err := rq.db.First(&res, "id=?", ID).Error; err != nil {
		return domain.Core{}, err
	}
	if err := rq.db.Delete(&res).Error; err != nil {
		return domain.Core{}, err
	}
	return ToDomain(res), nil
}

// func (rq *repoQuery) Delete() error {

// }
