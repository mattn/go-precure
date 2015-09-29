package precure

import (
	"sync"

	"github.com/mattn/go-precure/role"
)

var (
	m          sync.Mutex
	AllStars   []role.Precure
	AllBuddies []role.Buddy
)

func RegisterPrecure(p role.Precure) {
	m.Lock()
	defer m.Unlock()

	AllStars = append(AllStars, p)
}

func RegisterBuddy(b role.Buddy) {
	m.Lock()
	defer m.Unlock()

	AllBuddies = append(AllBuddies, b)
}
