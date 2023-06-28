package repository

import (
	"goAPI/pkg/models"

	"gorm.io/gorm"
)

// ItemRepository defines the methods that an item repository should have
type ItemRepository interface {
	GetItems() ([]models.Item, error)
	GetItemByID(id int) (*models.Item, error)
	AddItem(item models.Item) (*models.Item, error)
	RemoveItem(id int) error
	UpdateItem(item *models.Item) (*models.Item, error)
}

// ItemRepo Implements the interface
type ItemRepo struct {
	db *gorm.DB
}

// RemoveItem removes item from the database
func (r *ItemRepo) RemoveItem(id int) error {
	if err := r.db.Delete(&models.Item{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetItems returns a list of items from the database
func (r *ItemRepo) GetItems() ([]models.Item, error) {
	var items []models.Item
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// GetItemByID returns the item with the specified ID
func (r *ItemRepo) GetItemByID(id int) (*models.Item, error) {
	var item models.Item
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

// AddItem returns the item with the specified ID
func (r *ItemRepo) AddItem(item models.Item) (*models.Item, error) {
	if err := r.db.Create(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

// UpdateItem updates the existing item in the database
func (r *ItemRepo) UpdateItem(item *models.Item) (*models.Item, error) {
	if err := r.db.Save(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

// DI
func NewItemRepo(db *gorm.DB) *ItemRepo {
	return &ItemRepo{db: db}
}
