package lesson

import (
	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/element"
)

type UpdateEngine struct {
}

// Reports whether the element differs from its latest revision
func (e *UpdateEngine) HasDiff(el *element.Element, rev *element.Revision) bool {
	return el.ParentID() != rev.ParentID() ||
		!element.CompareConfigs(el, rev)
}

// Computes new revisions just for changed elements
func (e *UpdateEngine) ComputeRevisions(els []element.Element, revs []element.Revision) []element.Revision {
	res := make([]element.Revision, 0)
	cache := make(map[uuid.UUID]*element.Revision, len(revs))

	for i := range revs {
		rev := &revs[i]
		cache[rev.ElementID()] = rev
	}

	for i := range els {
		el := &els[i]
		rev := cache[el.ID()]

		if rev != nil && e.HasDiff(el, rev) {
			res = append(res, *el.ToRevision())
		}
	}

	return res
}
