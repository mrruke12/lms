package revision

import (
	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/element"
)

// Computes new revisions just for changed elements
func ComputeRevisions(els []element.Element, revs []Revision) []Revision {
	res := make([]Revision, 0)
	cache := make(map[uuid.UUID]*Revision, len(revs))

	for i := range revs {
		rev := &revs[i]
		cache[rev.ElementID()] = rev
	}

	for i := range els {
		el := &els[i]
		rev := cache[el.ID()]

		// if rev is nil then it's a new element so we must create revision, otherwise we create it if there's difference
		if rev == nil || rev.Differs(el) {
			res = append(res, *FromElement(el))
		}
	}

	return res
}
