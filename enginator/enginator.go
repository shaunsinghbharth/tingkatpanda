package enginator

import (
	"sync"
)

var (
	tables = make(map[string]*EnginatorTable)
	mutex sync.RWMutex
)

// Table returns the existing engine table with given name or creates a new one
// if the table does not exist yet.
func Table(table string) *EnginatorTable {
	mutex.RLock()
	t, ok := tables[table]
	mutex.RUnlock()

	if !ok {
		t = &EnginatorTable{
			name:  table,
			items: make(map[interface{}]*EnginatorItem),
		}

		mutex.Lock()
		tables[table] = t
		mutex.Unlock()
	}

	return t
}
