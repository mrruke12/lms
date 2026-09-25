package element

import (
	"github.com/google/uuid"
	"github.com/mrruke12/lms/internal/domains/element"
)

// Create a revision from element
func CreateFromElement(el *element.Element) *Revision {
	return &Revision{
		elementID:  el.ID(),
		lessonID:   el.LessonID(),
		parentID:   el.ParentID(),
		typ:        string(el.Type()),
		assessment: string(el.Assessment()),
		config:     el.ConfigRaw(),
	}
}

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
			res = append(res, *CreateFromElement(el))
		}
	}

	return res
}
