package eternal

import (
	"github.com/CharlesHolbrow/synk"
)

// ConstructContainer creates a container for an eternal synk Object
func ConstructContainer(typeKey string) synk.Object {
	switch typeKey {
	case "cell":
		return &Cell{}
	case "n":
		return &Note{}
	}
	return nil
}

// ConstructClient creates a custom client for the EternalApp
func ConstructClient(c synk.Client) synk.CustomClient {
	return Client{}
}
