package service

import (
	"context"
	"github.com/api_base/internal/domain"
	"github.com/api_base/internal/domain/model"
)

type Service interface {
	Get(ctx context.Context, id string) (*model.Model, error)
}

type service struct {
	container domain.Container
}

func NewService(container domain.Container) Service {
	return &service{
		container,
	}
}

func (s service) Get(ctx context.Context, id string) (*model.Model, error) {
	res, err := s.container.Repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
