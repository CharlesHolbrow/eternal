package eternal

import (
	"encoding/json"

	"github.com/CharlesHolbrow/synk"
)

// Cell represents a sound in the eternal music landscape
//@PA:cell
type Cell struct {
	synk.Tag  `bson:",inline"`
	AudioPath string  `json:"audioPath"`
	MapName   string  `json:"map"`
	Hue       float32 `json:"hue"`
	Class     string  `json:"cls"`
	X         int     `json:"x"`
	Y         int     `json:"y"`
	diff      cellDiff
}

func (n *Cell) String() string {
	json, _ := json.Marshal(n)
	return string(json)
}

// TypeKey identifies this type for synk Objects
func (n *Cell) TypeKey() string {
	return "cell"
}

// GetSubKey helps satisfy synk.Object
func (n *Cell) GetSubKey() string {
	cx, _ := split(n.GetX())
	cy, _ := split(n.GetY())
	return makeSubKey(n.GetMapName(), cx, cy)
}

// GetPrevSubKey helps satisfy synk.Object
func (n *Cell) GetPrevSubKey() string {
	cx, _ := split(n.GetPrevX())
	cy, _ := split(n.GetPrevY())
	return makeSubKey(n.GetPrevMapName(), cx, cy)
}
