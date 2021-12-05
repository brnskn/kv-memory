package store

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brnskn/kv-memory/internal/entity"
	"github.com/brnskn/kv-memory/pkg/response"
	"github.com/gorilla/mux"
)

func testApi(t *testing.T, method string, target string, f func(w http.ResponseWriter, r *http.Request)) string {
	if repo == nil {
		Init()
	}
	r := httptest.NewRequest(method, target, nil)
	w := httptest.NewRecorder()
	f(w, r)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %s", err.Error())
	}
	return strings.TrimSpace(string(data))
}

func TestApiHandlers(t *testing.T) {
	router := mux.NewRouter()
	RegisterHandlers(router)
}

func TestApiFlush(t *testing.T) {
	message := response.Response{
		Message: "store successfully flushed",
	}
	json, _ := json.Marshal(&message)
	want := string(json)
	got := testApi(t, http.MethodDelete, "/", Flush)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestApiSet(t *testing.T) {
	store := entity.Store{
		Key:   "foo",
		Value: "bar",
	}
	json, _ := json.Marshal(&store)
	want := string(json)
	got := testApi(t, http.MethodPost, "/?key=foo&value=bar", Set)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestApiGet(t *testing.T) {
	TestApiSet(t)
	store := entity.Store{
		Key:   "foo",
		Value: "bar",
	}
	json, _ := json.Marshal(&store)
	want := string(json)
	got := testApi(t, http.MethodGet, "/?key=foo", Get)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestApiGetError(t *testing.T) {
	message := response.Response{
		Message: "key not found in the store",
	}
	json, _ := json.Marshal(&message)
	want := string(json)
	got := testApi(t, http.MethodGet, "/?key=foo2", Get)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
