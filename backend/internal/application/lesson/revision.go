package lesson

import (
	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/element"
)

// Reports whether the element differs from its latest revision
func HasDiff(el *element.Element, rev *element.Revision) bool {
	return el.ParentID() != rev.ParentID() ||
		!element.CompareConfigs(el, rev)
}

// Computes new revisions just for changed elements
func ComputeRevisions(els []element.Element, revs []element.Revision) []element.Revision {
	res := make([]element.Revision, 0)
	cache := make(map[uuid.UUID]*element.Revision, len(revs))

	for i := range revs {
		rev := &revs[i]
		cache[rev.ElementID()] = rev
	}

	for i := range els {
		el := &els[i]
		rev := cache[el.ID()]

		// if rev is nil then it's a new element so we must create revision, otherwise we create it if there's difference
		if rev == nil || HasDiff(el, rev) {
			res = append(res, *el.ToRevision())
		}
	}

	return res
}
