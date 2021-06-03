package enginator

import (
	"sync"
	_ "time"
)

// Structure of an item in the recommendation engine.
// Parameter data contains the user-set value in the engine.
type EnginatorItem struct {
	sync.RWMutex

	// The item's key.
	key interface{}
	// All items for this key.
	data map[interface{}]float64
}

// CreateEnginatorItem returns a newly created EnginatorItem.
// Parameter key is the item's key.
// Parameter data is the item's value.
func CreateEnginatorItem(key interface{}, data map[interface{}]float64) EnginatorItem {
	return EnginatorItem{
		key:           key,
		data:          data,
	}
}

// Key returns the key of this item.
func (item *EnginatorItem) Key() interface{} {
	// immutable
	return item.key
}

// Data returns the value of this item.
func (item *EnginatorItem) Data() map[interface{}]float64 {
	// immutable
	return item.data
}
