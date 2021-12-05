package config

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	want := "bar"
	os.Setenv("foo", "bar")
	Load()
	got := Get("foo", "")
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestConfigDefaultValue(t *testing.T) {
	want := "bar"
	os.Setenv("foo", "bar")
	Load()
	got := Get("foo2", "bar")
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestConfigIntValue(t *testing.T) {
	want := 5
	os.Setenv("foo", "5")
	Load()
	got := GetInt("foo", 0)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestConfigIntDefaultValue(t *testing.T) {
	want := 5
	os.Setenv("foo", "5")
	Load()
	got := GetInt("foo2", 5)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestConfigIntDefaultValueError(t *testing.T) {
	want := 5
	os.Setenv("foo", "bar")
	Load()
	got := GetInt("foo", 5)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
