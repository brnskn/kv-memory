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

// Creates new repository
func NewRepository() *repository {
	return &repository{}
}

type repository struct{}

// Gets value of given key from the store and returns the store entity
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

// Sets given value to given key on the store and returns the store entity
func (repository *repository) Set(key string, value string) entity.Store {
	store.Instance().Set(key, value)
	return entity.Store{
		Key:   key,
		Value: value,
	}
}

// Deletes all items from the store
func (repository *repository) Flush() {
	store.Instance().Flush()
}
