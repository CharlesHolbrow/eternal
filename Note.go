package eternal

import (
	"encoding/json"

	"github.com/CharlesHolbrow/synk"
)

// Note is a minimal example of a synk.Object. Note that we use pagen to
// generate the required methods.
//@PA:n
type Note struct {
	synk.Tag `bson:",inline"`
	SubKey   string `json:"subKey"`
	Number   int    `json:"number"`
	Velocity int    `json:"velocity"`
	diff     noteDiff
}

func (n *Note) String() string {
	json, _ := json.Marshal(n)
	return string(json)
}

// TypeKey identifies this type for synk Objects
func (n *Note) TypeKey() string {
	return "n"
}
