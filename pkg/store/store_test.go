package store

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestStoreSingleton(t *testing.T) {
	store1 := Instance()
	store2 := Instance()
	if store1 != store2 {
		t.Errorf("Instance() function is not returning single instance of Store")
	}
}

func TestStoreGet(t *testing.T) {
	want := "bar"
	store := new()
	store.Set("foo", "bar")
	got, _ := store.Get("foo")
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestStoreFlush(t *testing.T) {
	want := ""
	store := new()
	store.Set("foo", "bar")
	store.Flush()
	got, _ := store.Get("foo")
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestStoreSave(t *testing.T) {
	want := fmt.Sprint(rand.Intn(100))
	{
		store := new()
		store.Set("foo", want)
		store.Save()
	}
	store := new()
	got, _ := store.Get("foo")
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestStoreAutoSave(t *testing.T) {
	backups, err := filepath.Glob("/tmp/*-data.json")
	if err == nil {
		for _, backup := range backups {
			os.Remove(backup)
		}
	}
	want := fmt.Sprint(rand.Intn(100))
	{
		store := new()
		store.StartAutoSaver(time.Second)
		store.Set("foo", want)
		time.Sleep(2 * time.Second)
		store.StopAutoSaver()
	}
	store := new()
	got, _ := store.Get("foo")
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
