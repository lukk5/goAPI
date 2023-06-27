//go:build wireinject
// +build wireinject

package extensions

import (
	"github.com/google/wire"
	"goAPI/pkg/context"
	"goAPI/pkg/handlers"
	"goAPI/pkg/repository"
)

func InitializeItemHandler() *handlers.ItemHandler {
	wire.Build(handlers.NewItemHandler, repository.NewItemRepo, context.NewDbConnection, RepoAsRepository)
	return &handlers.ItemHandler{}
}

// RepoAsRepository is a Wire provider function that provides an ItemRepository from an ItemRepo.
func RepoAsRepository(repo *repository.ItemRepo) repository.ItemRepository {
	return repo
}
