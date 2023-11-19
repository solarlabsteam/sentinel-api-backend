package event

import (
	"strconv"

	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/solarlabsteam/sentinel-api-backend/types"
)

func GetSessionEventStart(items []abcitypes.Event) (*types.Event, error) {
	return types.NewEventFromABCIEvents(items, "sentinel.session.v2.EventStart")
}

func GetSessionIDFromABCIEvents(items []abcitypes.Event) (uint64, error) {
	item, err := GetSessionEventStart(items)
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(item.Attributes["id"], 10, 64)
}
