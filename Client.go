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
func (cc Client) OnConnect(client *synk.Client) {
	log.Println("Custom Client Connected:", client.ID())
}

// OnMessage is called when the client sends a message
func (cc Client) OnMessage(client *synk.Client, method string, data []byte) {
	log.Println("Custom Client Message:", method)
	switch method {
	case "note":
		note := &NoteEvent{}
		if err := json.Unmarshal(data, note); err == nil {
			fmt.Println("Note:", note)
			cc.noteEvent(client, *note)
		} else {
			fmt.Println("Client send bad note event:", err)
		}
	}
}

// OnSubscribe is called with the client changes their subscription
func (cc Client) OnSubscribe(client *synk.Client, subKeys []string, objs []synk.Object) {
	log.Printf("Custom Client: Subscription add(%d) objs(%d)", len(subKeys), len(objs))
}

func (cc Client) noteEvent(client *synk.Client, ne NoteEvent) {
	if json, err := json.Marshal(ne); err == nil {
		client.Loader.Publish("piano", json)
	} else {
		fmt.Println("eternal client failed to marshall NoteEvent json", err.Error())
	}
}

/***************************************************************
/
/ Messages that may be sent from clients
/
***************************************************************/

// NoteEvent messages may be send by the client
type NoteEvent struct {
	N  int
	V  int
	C  int
	On bool
}

func (n NoteEvent) String() string {
	return fmt.Sprintf("Note:%d vel:%d chan:%d on:%t", n.N, n.V, n.C, n.On)
}
