package types

import (
	"fmt"

	abcitypes "github.com/tendermint/tendermint/abci/types"
)

type Event struct {
	Type       string            `json:"type,omitempty" bson:"type"`
	Attributes map[string]string `json:"attributes,omitempty" bson:"attributes"`
}

func NewEventFromABCIEvent(v *abcitypes.Event) *Event {
	item := &Event{
		Type:       v.Type,
		Attributes: make(map[string]string),
	}

	for _, x := range v.Attributes {
		vLen := len(x.Value)
		if vLen >= 2 {
			if x.Value[0] == '"' && x.Value[vLen-1] == '"' {
				x.Value = x.Value[1 : vLen-1]
			}
		}

		item.Attributes[string(x.Key)] = string(x.Value)
	}

	return item
}

func NewEventFromABCIEvents(items []abcitypes.Event, v string) (*Event, error) {
	for i := 0; i < len(items); i++ {
		if items[i].Type == v {
			return NewEventFromABCIEvent(&items[i]), nil
		}
	}

	return nil, fmt.Errorf("event %s does not exist", v)
}
