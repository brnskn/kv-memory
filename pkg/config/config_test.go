package config_test

import (
	"os"
	"testing"

	"github.com/brnskn/kv-memory/pkg/config"
)

func TestConfig(t *testing.T) {
	want := "bar"
	os.Setenv("foo", "bar")
	config.Load()
	got := config.Get("foo", "")
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestConfigDefaultValue(t *testing.T) {
	want := "bar"
	os.Setenv("foo", "bar")
	config.Load()
	got := config.Get("foo2", "bar")
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestConfigIntValue(t *testing.T) {
	want := 5
	os.Setenv("foo", "5")
	config.Load()
	got := config.GetInt("foo", 0)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestConfigIntDefaultValue(t *testing.T) {
	want := 5
	os.Setenv("foo", "5")
	config.Load()
	got := config.GetInt("foo2", 5)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
