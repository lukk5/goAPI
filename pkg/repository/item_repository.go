package repository

import (
	"database/sql"
	"fmt"
	"github.com/lukk5/goAPI/pkg/models"
)

// ItemRepository defines the methods that an item repository should have
type ItemRepository interface {
	GetItems() ([]models.Item, error)
	GetItemByID(id int) (*models.Item, error)
	AddItem(item models.Item) (*models.Item, error)
}

// ItemRepo Implements the interface
type ItemRepo struct {
	db *sql.DB
}

// GetItems returns a list of items from the in-memory repository
func (r *ItemRepo) GetItems() ([]models.Item, error) {
	rows, err := r.db.Query(`SELECT "ID", "Name"
	FROM public."Item";`)
	if err != nil {
		return nil, fmt.Errorf("error getting items: %v", err)
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %v", err)
	}

	return items, nil

}

// GetItemByID returns the item with the specified ID
func (r *ItemRepo) GetItemByID(id int) (*models.Item, error) {
	row := r.db.QueryRow(`SELECT "ID", "Name"
	FROM public."Item" WHERE "ID" = $1;`, id)
	var item models.Item
	switch err := row.Scan(&item.ID, &item.Name); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("item was not found")
	case nil:
		return &item, nil
	default:
		return nil, fmt.Errorf("error scanning row: %v", err)
	}
}

// AddItem returns the item with the specified ID
func (r *ItemRepo) AddItem(item models.Item) (*models.Item, error) {
	sqlStatement := `INSERT INTO public."Item"("Name")
	VALUES ($1) RETURNING "ID";`

	id := 0
	err := r.db.QueryRow(sqlStatement, item.Name).Scan(&id)
	if err != nil {
		return nil, err
	}

	item.ID = id
	return &item, nil
}

// DI
func NewItemRepo(db *sql.DB) *ItemRepo {
	return &ItemRepo{db: db}
}
