package handlers

import (
	"goAPI/pkg/models"
	"net/http"
	"strconv"

	"goAPI/pkg/repository"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	Repo repository.ItemRepository
}

// AddItemHandler @Summary Add a new item
// @Description Add item to the database
// @ID add-item
// @Accept  json
// @Produce  json
// @Param item body models.Item true "Add item"
// @Success 200 {object} models.Item
// @Router /items [post]
func (h *ItemHandler) AddItemHandler(c *gin.Context) error {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		return err
	}

	item, err := h.Repo.AddItem(newItem)
	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, item)
	return nil
}

// GetItemsHandler @Summary Get all items
// @Description Get all items from the database
// @ID get-items
// @Produce  json
// @Success 200 {array} models.Item
// @Router /items [
func (h *ItemHandler) GetItemsHandler(c *gin.Context) error {
	items, err := h.Repo.GetItems()
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, items)
	return nil
}

// GetItemByIdHandler @Summary Get an item by ID
// @Description Get a specific item from the database
// @ID get-item
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {object} models.Item
// @Router /items/{id} [get]
func (h *ItemHandler) GetItemByIdHandler(c *gin.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}

	item, err := h.Repo.GetItemByID(id)

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, item)
	return nil
}

// NewItemHandler DI
func NewItemHandler(repo repository.ItemRepository) *ItemHandler {
	return &ItemHandler{repo}
}
