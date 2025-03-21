package core

import (
	"sync"
)

var (
	running bool = false
	mutex   sync.Mutex
)

func IsRunning() bool {
	mutex.Lock()
	defer mutex.Unlock()
	return running
}

func SetRunning(value bool) {
	mutex.Lock()
	defer mutex.Unlock()
	running = value
}

func ToggleRunning() {
	mutex.Lock()
	defer mutex.Unlock()
	running = !running
}
