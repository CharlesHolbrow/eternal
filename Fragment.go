package eternal

import (
	"github.com/CharlesHolbrow/synk"
)

// Fragment stores the contents of a subscription key
type Fragment struct {
	Mutator synk.Mutator
	sKey1   string
	On      [128]*Note
}

// NewFragment - create a Fragment
//
// Requires a clean Mutator
func NewFragment(k1 string, mutator synk.Mutator) *Fragment {

	notes := &Fragment{
		sKey1:   k1,
		Mutator: mutator,
	}

	objects, err := notes.Mutator.Load([]string{k1})
	if err != nil {
		panic("Error initializing eternal Fragment: " + err.Error())
	}

	for _, obj := range objects {
		if obj, ok := obj.(*Note); ok {
			if obj.Number < 0 || obj.Number > 127 {
				mutator.Delete(obj)
				continue
			}
			if notes.On[obj.Number] != nil {
				mutator.Delete(obj)
				continue
			}
			notes.On[obj.Number] = obj
		}
	}

	return notes
}

// AddNote to the Fragment. The note's subscription key will be set.
// No-op if we already have a note with n.Number
func (frag *Fragment) AddNote(n *Note) {
	// Ensure SubKey
	if n.GetSubKey() == "" {
		n.SetSubKey(frag.sKey1)
	}

	if frag.On[n.Number] != nil {
		return
	}

	frag.Mutator.Create(n) // Ensures ID
	frag.On[n.Number] = n
}

// RemoveNote removes any note in the fragment with number equal to n.Number
func (frag *Fragment) RemoveNote(n *Note) {
	if n.Number < 0 || n.Number > 127 {
		return
	}

	removeMe := frag.On[n.Number]
	if removeMe == nil {
		return
	}

	frag.Mutator.Delete(removeMe)
	frag.On[n.Number] = nil
}

// RemoveAllNotes is self explanitory
func (frag *Fragment) RemoveAllNotes() {
	for id, note := range frag.On {
		if note == nil {
			continue
		}
		frag.Mutator.Delete(note)
		frag.On[id] = nil
	}
}

// K1 Gets the first subscription key
func (frag *Fragment) K1() string {
	return frag.sKey1
}
