// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package extensions

import (
	"github.com/lukk5/goApi/pkg/handlers"
	"github.com/lukk5/goApi/pkg/repository"
)

// Injectors from wire.go:

func InitializeItemHandler() *handlers.ItemHandler {
	db := repository.NewDbConnection()
	itemRepo := repository.NewItemRepo(db)
	itemRepository := RepoAsRepository(itemRepo)
	itemHandler := handlers.NewItemHandler(itemRepository)
	return itemHandler
}

// wire.go:

// RepoAsRepository is a Wire provider function that provides an ItemRepository from an ItemRepo.
func RepoAsRepository(repo *repository.ItemRepo) repository.ItemRepository {
	return repo
}
