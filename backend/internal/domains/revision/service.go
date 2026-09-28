package revision

import (
	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/element"
)

// Computes new revisions just for changed elements.
//
//	return updatedEls, newRevs
func ComputeRevisions(els []element.Element, revs []Revision) ([]element.Element, []Revision) {
	var newRevs []Revision
	var updatedEls []element.Element

	cache := make(map[uuid.UUID]*Revision, len(revs))

	for i := range revs {
		rev := &revs[i]
		cache[rev.ElementID()] = rev
	}

	for i := range els {
		el := &els[i]
		rev, exists := cache[el.ID()]

		// if rev is nil then it's a new element so we must create revision, otherwise we create it if there's difference
		if !exists || rev.Differs(el) {
			newRevs = append(newRevs, *FromElement(el))
			updatedEls = append(updatedEls, *el)
		}
	}

	return updatedEls, newRevs
}
