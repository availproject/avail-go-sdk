package examples

import (
	"fmt"
	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func Run_events() {
	sdk := SDK.NewSDK(SDK.LocalEndpoint)

	acc, err := SDK.Account.Alice()
	if err != nil {
		panic(err)
	}

	tx := sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(1))
	if err != nil {
		panic(err)
	}

	eventsM := res.Events
	if eventsM.IsNone() {
		panic("Failed to decode events")
	}
	events := eventsM.Unwrap()
	for _, event := range events {
		println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v`, event.PalletName, event.PalletIndex, event.EventName, event.EventIndex))
	}

	// EventFindAll, EventFindFirst, EventFindLast
	eventM2 := SDK.EventFindFirst(events, daPallet.EventDataSubmitted{})
	if eventM2.IsNone() {
		panic("No event with that description was found or one was found but we failed to decode it")
	}
	event2 := eventM2.Unwrap()
	println(fmt.Sprintf(`Who: %v, Datahash: %v`, event2.Who.ToHuman(), event2.DataHash.ToHexWith0x()))

	// EventFindAllChecked, EventFindFirstChecked, EventFindLastChecked
	eventM3, err := SDK.EventFindFirstChecked(events, daPallet.EventDataSubmitted{})
	if err != nil {
		panic("Found the even but failed to decode it")
	}
	if eventM3.IsNone() {
		panic("No event with that description was found")
	}
	event3 := eventM3.Unwrap()
	println(fmt.Sprintf(`Who: %v, Datahash: %v`, event3.Who.ToHuman(), event3.DataHash.ToHexWith0x()))

}
