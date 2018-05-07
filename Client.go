package eternal

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/CharlesHolbrow/synk"
)

// Client implements custom handlers
type Client struct {
}

// OnConnect is called when client connects via WebSocket
func (cc Client) OnConnect(client synk.Client) {
	log.Println("Custom Client Connected:", client.ID())
}

// OnMessage is called when the client sends a message
func (cc Client) OnMessage(client synk.Client, method string, data []byte) {
	log.Println("Custom Client Message:", method)
	switch method {
	case "moveCell":
		event := &MoveCellEvent{}
		if err := json.Unmarshal(data, event); err == nil {
			log.Println("MoveCellEvent:", event)
			cc.handleMoveCellEvent(client, *event)
		} else {
			fmt.Println("CLient requested bad moveCell event:", err)
		}
	}
}

// OnSubscribe is called with the client changes their subscription
func (cc Client) OnSubscribe(client synk.Client, subKeys []string, objs []synk.Object) {
	log.Printf("Custom Client: Subscription add(%d) objs(%d)", len(subKeys), len(objs))
}

func (cc Client) handleMoveCellEvent(client synk.Client, e MoveCellEvent) {
	if json, err := json.Marshal(e); err == nil {
		client.Publish("piano", json)
	} else {
		log.Println("failed to re-marshall MoveCellEvent")
	}
}

/***************************************************************
/
/ Messages that may be sent from clients
/
***************************************************************/

// MoveCellEvent is send by the client as a request
type MoveCellEvent struct {
	ID string
	X  int
	Y  int
}

func (e MoveCellEvent) String() string {
	return fmt.Sprintf("Move %s to (%d, %d)", e.ID, e.X, e.Y)
}
