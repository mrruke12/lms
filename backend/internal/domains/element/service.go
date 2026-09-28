package element

import (
	"github.com/google/uuid"
)

// Categorizes elements into persisted, new, deleted i.e. which to update, create, delete in the storage.
//
//	return persisted, new, deleted
func Categorize(memory, storage []Element) ([]Element, []Element, []Element) {
	var new []Element
	var deleted []Element
	var persisted []Element

	memoryCache := make(map[uuid.UUID]struct{})
	storageCache := make(map[uuid.UUID]struct{})

	for i := range memory {
		id := memory[i].id
		memoryCache[id] = struct{}{}
	}

	for i := range storage {
		id := storage[i].id
		storageCache[id] = struct{}{}
	}

	for _, el := range memory {
		if _, exists := storageCache[el.id]; !exists {
			new = append(new, el)
		} else {
			persisted = append(persisted, el)
		}
	}

	for _, el := range storage {
		if _, exists := memoryCache[el.id]; !exists {
			deleted = append(deleted, el)
		}
	}

	return persisted, new, deleted
}
