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
func (h *ItemHandler) AddItemHandler(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		panic(err)
	}

	item, err := h.Repo.AddItem(newItem)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, item)
}

// GetItemsHandler @Summary Get all items
// @Description Get all items from the database
// @ID get-items
// @Produce  json
// @Success 200 {array} models.Item
// @Router /items [
func (h *ItemHandler) GetItemsHandler(c *gin.Context) {
	items, err := h.Repo.GetItems()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, items)
}

// GetItemByIdHandler @Summary Get an item by ID
// @Description Get a specific item from the database
// @ID get-item
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {object} models.Item
// @Router /items/{id} [get]
func (h *ItemHandler) GetItemByIdHandler(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		panic(err)
	}

	item, err := h.Repo.GetItemByID(id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, item)
}

// NewItemHandler DI
func NewItemHandler(repo repository.ItemRepository) *ItemHandler {
	return &ItemHandler{repo}
}
