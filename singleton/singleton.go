package singleton

import (
	"fmt"
	"sync"
)

type repository struct {
	configs map[string]string
	mu    sync.RWMutex
}

func (r *repository) Set(key, data string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.configs[key] = data
}

func (r *repository) Get(key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	config, ok := r.configs[key]
	if !ok {
		return "", fmt.Errorf("The '%s' is not presented", key)
	}
	return config, nil
}

var (
	r    *repository
	once sync.Once
)

func ConfigRepository() *repository {
	once.Do(func() {

		fmt.Printf("Initializing singleton object\n")

		// can be reading configs from a source
		_configs := make(map[string]string)
		_configs["Config 1"] = "Value 1"
		_configs["Config 2"] = "Value 2"


		r = &repository{
			configs: _configs,
		}
	})

	return r
}

