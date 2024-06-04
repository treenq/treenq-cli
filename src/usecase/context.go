package usecase

import (
	"context"

	"github.com/treenq/treenq-cli/src/models"
)

type Store interface {
	SetActiveContext(name string) error
	GetContexts() []models.Context
	NewContext(ctx models.Context) error
}

type ContextUsecase struct {
	store Store
}

func NewContextUsecase(store Store) *ContextUsecase {
	return &ContextUsecase{
		store: store,
	}
}

// NewContext crates new context and sets it as active
func (u *ContextUsecase) NewContext(ctx context.Context, name, url string) error {
	return u.store.NewContext(models.Context{
		Name: name,
		Url:  url,
	})
}

func (u *ContextUsecase) SetContext(ctx context.Context, name string) error {
	return u.store.SetActiveContext(name)
}

func (u *ContextUsecase) ListContexts(ctx context.Context) ([]models.Context, error) {
	return u.store.GetContexts(), nil
}
