package singleton

import (
	"sync"
)

type Singleton map[string]int

var once sync.Once
var instance Singleton

func New() Singleton  {
	once.Do(func() {
			instance = make(Singleton)
	})
	return instance
}
