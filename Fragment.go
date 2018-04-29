package eternal

import (
	"fmt"
	"image"
	"log"

	"github.com/CharlesHolbrow/synk"
)

// Fragment stores the contents of a subscription key
type Fragment struct {
	Mutator synk.Mutator
	MapName string
	On      [128]*Note
	Cells   map[string]*Cell
	Keys    []string
}

// NewFragment - create a Fragment
//
// Requires a clean Mutator
func NewFragment(mapName string, area image.Rectangle, mutator synk.Mutator) *Fragment {

	size := area.Size()
	chunkCount := size.X * size.Y
	chunkKeys := make([]string, 0, chunkCount)

	for y := area.Min.Y; y < area.Max.Y; y++ {
		for x := area.Min.X; x < area.Max.X; x++ {
			chunkKeys = append(chunkKeys, makeSubKey(mapName, x, y))
		}
	}

	frag := &Fragment{
		Keys:    chunkKeys,
		MapName: mapName,
		Mutator: mutator,
		Cells:   make(map[string]*Cell),
	}

	objects, err := frag.Mutator.Load(frag.Keys)
	if err != nil {
		panic("Error initializing eternal Fragment: " + err.Error())
	}

	for _, v := range objects {

		switch obj := v.(type) {
		case *Cell:
			frag.Cells[obj.TagID] = obj
		case *Note:
			if obj.Number < 0 || obj.Number > 127 {
				mutator.Delete(obj)
				continue
			}
			if frag.On[obj.Number] != nil {
				mutator.Delete(obj)
				continue
			}
			frag.On[obj.Number] = obj
		}
	}

	fmt.Printf("Found %d Cells\n", len(frag.Cells))

	return frag
}

// AddCell to the Fragment
func (frag *Fragment) AddCell(c *Cell) {
	c.SetMapName(frag.MapName)
	frag.Mutator.Create(c) // Ensures TagID
	frag.Cells[c.TagID] = c
}

// RemoveCell from the fragment
func (frag *Fragment) RemoveCell(c *Cell) {
	err := frag.Mutator.Delete(c)
	if err != nil {
		panic("error deleting cell with ID: " + c.TagID)
	}
	delete(frag.Cells, c.TagID)
}

// AddNote to the Fragment. The note's subscription key will be set.
// No-op if we already have a note with n.Number
func (frag *Fragment) AddNote(n *Note) {
	// Ensure SubKey
	if n.GetSubKey() == "" {
		log.Println("you must now spceify a subkey when adding a Note")
		return
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
