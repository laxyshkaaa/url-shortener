package link

import (
	"gorm.io/gorm/clause"
	"url/pkg/db"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}
func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}
func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}
func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (repo *LinkRepository) GetByID(id uint) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&Link{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}
func (repo *LinkRepository) GetLinks(limit, offset int) []Link {
	var links []Link
	repo.Database.Table("links").Where("deleted_at is null").Limit(limit).Offset(offset).Scan(&links)
	return links
}

func (repo *LinkRepository) Count() int64 {
	var count int64
	repo.Database.Table("links").Where("deleted_at is null").Count(&count)
	return count
}
