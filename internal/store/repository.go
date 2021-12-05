package store

import (
	"errors"

	"github.com/brnskn/kv-memory/internal/entity"
	"github.com/brnskn/kv-memory/pkg/store"
)

// Repository encapsulates the logic to access stored values.
type Repository interface {
	Get(key string) (entity.Store, error)
	Set(key string, value string) entity.Store
	Flush()
}

var _ Repository = &repository{}

func NewRepository() *repository {
	return &repository{}
}

type repository struct{}

func (repository *repository) Get(key string) (entity.Store, error) {
	value, found := store.Instance().Get(key)
	if !found {
		return entity.Store{}, errors.New("key not found in the store")
	}
	return entity.Store{
		Key:   key,
		Value: value,
	}, nil
}

func (repository *repository) Set(key string, value string) entity.Store {
	store.Instance().Set(key, value)
	return entity.Store{
		Key:   key,
		Value: value,
	}
}

func (repository *repository) Flush() {
	store.Instance().Flush()
}
