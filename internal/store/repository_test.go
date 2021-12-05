package store

import (
	"reflect"
	"testing"

	"github.com/brnskn/kv-memory/internal/entity"
)

func TestRepoSet(t *testing.T) {
	want := entity.Store{
		Key:   "foo",
		Value: "bar",
	}
	got := NewRepository().Set(want.Key, want.Value)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRepoGet(t *testing.T) {
	want := entity.Store{
		Key:   "foo",
		Value: "bar",
	}
	got, err := NewRepository().Get(want.Key)
	if err != nil {
		t.Errorf("expected error to be nil got %s", err.Error())
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRepoFlush(t *testing.T) {
	NewRepository().Flush()
	store := entity.Store{
		Key:   "foo",
		Value: "bar",
	}
	want := "key not found in the store"
	_, err := NewRepository().Get(store.Key)
	if err == nil {
		t.Errorf("expected error to be not nil")
	}
	if err.Error() != want {
		t.Errorf("got %v want %v", err.Error(), want)
	}
}
