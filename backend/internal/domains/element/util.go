package element

import "slices"

// Returns true if configs are equal.
// Does not copy slices underhood
func CompareConfigs(el *Element, rev *Revision) bool {
	return slices.Equal(el.config, rev.config)
}
