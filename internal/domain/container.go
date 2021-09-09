package domain

import (
	"context"
	"github.com/api_base/config"
	"github.com/api_base/internal/domain/model"
	"github.com/api_base/internal/repository"
	"github.com/api_base/tool/database"
	"log"
)

type Container struct {
	Repo Repository
}

type Repository interface {
	Get(ctx context.Context, id string) (*model.Model, error)
}

func NewContainer(config config.Config) Container {
	db, err := database.NewRepository(config.Database)
	if err != nil {
		log.Fatal("initialize database fail: ", err)
	}
	return Container{
		Repo: repository.NewRepository(db),
	}
}
