package eternal

import (
	"github.com/CharlesHolbrow/synk"
)

// ConstructContainer creates a container for a eternal synk Object
func ConstructContainer(typeKey string) synk.Object {
	switch typeKey {
	case "n":
		return &Note{}
	}
	return nil
}

// NewNode creates a new synk node. This node may be a mutator or http handler
func NewNode() *synk.Node {
	node := synk.NewNode()
	node.NewContainer = ConstructContainer
	node.NewClient = func(client *synk.Client) synk.CustomClient {
		return Client{}
	}
	return node
}
