// This package provides an in-memory key-value store with a singleton pattern
// to ensure there is only one instance created by its used packages.
package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type Store struct {
	items     map[string]string
	mutex     *sync.RWMutex // To make store's read&write operations thread-safe
	autoSaver *autoSaver
}

var (
	once     sync.Once
	instance *Store
)

// Returns single instance of Store
func Instance() *Store {
	once.Do(func() {
		instance = new()
	})
	return instance
}

func new() *Store {
	store := &Store{
		items: make(map[string]string),
		mutex: &sync.RWMutex{},
	}
	store.loadLastBackup()
	return store
}

// Starts auto saver with given interval if the interval
// is not greater than zero it won't start auto saver.
func (store *Store) StartAutoSaver(autoSaveInterval time.Duration) {
	if autoSaveInterval > 0 {
		runAutoSaver(store, autoSaveInterval)
		// Auto saver goroutine will be automatically
		// stopped when the object is garbage-collected
		runtime.SetFinalizer(store, stopAutoSaver)
	}
}

// Stops the auto saver goroutine
func (store *Store) StopAutoSaver() {
	stopAutoSaver(store)
}

// Sets the given value to the given key
func (store *Store) Set(key string, value string) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.items[key] = value
}

// Gets the value of the given key
func (store *Store) Get(key string) (string, bool) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	item, found := store.items[key]
	if !found {
		return "", false
	}
	return item, true
}

// Deletes all items from the key-value store
func (store *Store) Flush() {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.items = make(map[string]string)
}

// Manually saves all items to the disk. The last saved item will
// be automatically restored when the Store object is regenerated.
func (store *Store) Save() {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	data, err := json.Marshal(store.items)
	if err == nil {
		path := fmt.Sprintf("/tmp/%d-data.json", time.Now().Unix())
		ioutil.WriteFile(path, data, 0700)
	}
}

func (store *Store) lastBackup() (string, bool) {
	backups, err := filepath.Glob("/tmp/*-data.json")
	if err == nil && len(backups) > 0 {
		lastBackup := backups[len(backups)-1]
		return lastBackup, true
	}
	return "", false
}

func (store *Store) loadLastBackup() {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	lastBackup, found := store.lastBackup()
	if found {
		file, err := ioutil.ReadFile(lastBackup)
		if err == nil {
			json.Unmarshal(file, &store.items)
		}
	}
}

type autoSaver struct {
	Interval time.Duration
	stop     chan bool
}

func (a *autoSaver) run(store *Store) {
	ticker := time.NewTicker(a.Interval)
	for {
		select {
		case <-ticker.C:
			store.Save()
		case <-a.stop:
			ticker.Stop()
			return
		}
	}
}

func stopAutoSaver(store *Store) {
	if store.autoSaver != nil {
		store.autoSaver.stop <- true
	}
}

func runAutoSaver(store *Store, interval time.Duration) {
	a := &autoSaver{
		Interval: interval,
		stop:     make(chan bool),
	}
	store.autoSaver = a
	go a.run(store)
}
