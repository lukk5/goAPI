//go:build wireinject
// +build wireinject

package extensions

import (
	"github.com/google/wire"
	"github.com/lukk5/goApi/pkg/handlers"
	"github.com/lukk5/goApi/pkg/repository"
)

func InitializeItemHandler() *handlers.ItemHandler {
	wire.Build(handlers.NewItemHandler, repository.NewItemRepo, repository.NewDbConnection, RepoAsRepository)
	return &handlers.ItemHandler{}
}

// RepoAsRepository is a Wire provider function that provides an ItemRepository from an ItemRepo.
func RepoAsRepository(repo *repository.ItemRepo) repository.ItemRepository {
	return repo
}
