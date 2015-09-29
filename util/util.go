package util

import (
	"fmt"
	"sync"
)

var (
	m sync.Mutex
)

func Say(s string) {
	m.Lock()
	defer m.Unlock()
	fmt.Println(s)
}
