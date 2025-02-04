package sdk

import (
	"errors"
	"fmt"
	"slices"

	"github.com/availproject/avail-go-sdk/interfaces"
	meta "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type Events struct {
	eventBytes []byte
	metadata   *meta.Metadata
	startIdx   uint32 /// Start index for events. Here we ignore the event count bytes
	eventCount uint32
}

func NewEvents(eventBytes []byte, metadata *meta.Metadata) (Events, error) {
	events := Events{}

	// Sanity Check
	if len(eventBytes) == 0 {
		return Events{}, nil
	}

	// Decode EventCount
	decoder := prim.NewDecoder(eventBytes, 0)
	compactEventCount := prim.CompactU32{}
	if err := decoder.Decode(&compactEventCount); err != nil {
		return Events{}, newError(err, ErrorCode004)
	}

	// Sanity Check. A stupid one but better than nothing
	if events.eventCount > 1_000_000 {
		return events, errors.New("Cannot decode events. Read garbage for event count")
	}

	events.eventCount = compactEventCount.Value
	events.startIdx = uint32(decoder.ScaleBytes.Offset)
	events.eventBytes = eventBytes
	events.metadata = metadata

	return events, nil
}

func (this *Events) Decode() (EventRecords, error) {
	result := EventRecords{}
	position := 0

	if this.eventCount == 0 {
		return EventRecords{}, nil
	}

	decoder := prim.NewDecoder(this.eventBytes, int(this.startIdx))
	for {
		eventRecord, err := NewEventRecord(&decoder, uint32(position), this.metadata)
		if err != nil {
			return EventRecords{}, err
		}

		result = append(result, eventRecord)
		position += 1

		if position == int(this.eventCount) {
			break
		}
	}

	// Sanity Check
	if decoder.RemainingLength() != 0 {
		return EventRecords{}, errors.New("All events were decoded but some bytes are left. This is not good.")
	}

	return result, nil
}

type EventPhase struct {
	VariantIndex   uint8
	ApplyExtrinsic prim.Option[uint32]
}

func (this *EventPhase) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "ApplyExtrinsic"
	case 1:
		return "Finalization"
	case 2:
		return "Initialization"
	default:
		panic("Unknown EventPhase Variant Index")
	}
}

func DecodeEventPhase(decoder *prim.Decoder) (EventPhase, error) {
	var eventPhase = EventPhase{}
	if err := decoder.Decode(&eventPhase.VariantIndex); err != nil {
		return EventPhase{}, newError(err, ErrorCode004)
	}

	switch eventPhase.VariantIndex {
	case 0:
		var value = uint32(0)
		if err := decoder.Decode(&value); err != nil {
			return EventPhase{}, newError(err, ErrorCode004)
		}
		eventPhase.ApplyExtrinsic.Set(value)
		return eventPhase, nil
	case 1:
		return eventPhase, nil
	case 2:
		return eventPhase, nil
	default:
		return eventPhase, errors.New(fmt.Sprintf(`Failed to Decode Event Phase. Unknown Variant Index %v`, eventPhase.VariantIndex))
	}
}

type EventRecords = []EventRecord

type EventRecord struct {
	Phase       EventPhase
	PalletIndex uint8
	EventIndex  uint8
	PalletName  string
	EventName   string
	Position    uint32
	AllBytes    []byte
	/// Phase, pallet/event index, fields and topic
	StartIdx uint32
	/// (After Phase) pallet/event index, fields and topic
	EventStartIndex uint32
	/// (After Phase, pallet/event index) fields and topic
	EventFieldsStartIndex uint32
	/// End Of the fields (before topic)
	EventFieldsEndIndex uint32
	/// end of everything
	EndIdx   uint32
	Metadata *meta.Metadata
	Topics   []prim.H256
}

func (this *EventRecord) TxIndex() prim.Option[uint32] {
	if this.Phase.ApplyExtrinsic.IsNone() {
		return prim.NewNone[uint32]()
	}

	return prim.NewSome(this.Phase.ApplyExtrinsic.Unwrap())
}

func NewEventRecord(decoder *prim.Decoder, position uint32, metadata *meta.Metadata) (EventRecord, error) {
	var eventRecord = EventRecord{}

	eventRecord.StartIdx = uint32(decoder.Offset())

	/* 	fmt.Println("Before Event Phase") */
	eventPhase, err := DecodeEventPhase(decoder)
	if err != nil {
		return EventRecord{}, err
	}
	eventRecord.Phase = eventPhase
	/* 	fmt.Println("Before After Phase") */

	// Pallet and Event Index Decoding

	eventRecord.EventStartIndex = uint32(decoder.Offset())
	if err := decoder.Decode(&eventRecord.PalletIndex); err != nil {
		return EventRecord{}, newError(err, ErrorCode004)
	}
	if err := decoder.Decode(&eventRecord.EventIndex); err != nil {
		return EventRecord{}, newError(err, ErrorCode004)
	}

	// Decoding Pallet and Event Names
	palletName, eventName, err := metadata.PalletEventName(eventRecord.PalletIndex, eventRecord.EventIndex)
	if err != nil {
		return EventRecord{}, err
	}
	eventRecord.PalletName = palletName
	eventRecord.EventName = eventName

	// fmt.Println(fmt.Sprintf(`Decoding %v, %v %v`, eventRecord.PalletName, eventRecord.EventName, eventRecord.Phase.ToString()))

	// Decode Fields
	eventRecord.EventFieldsStartIndex = uint32(decoder.Offset())
	if err := metadata.DecodeEvent(eventRecord.PalletIndex, eventRecord.EventIndex, decoder); err != nil {
		return EventRecord{}, err
	}
	eventRecord.EventFieldsEndIndex = uint32(decoder.Offset())
	/* 	fmt.Println("Before topics")

	   	fmt.Println(fmt.Sprintf(`EventFieldsStartIndexed %v`, eventRecord.EventFieldsStartIndex))
	   	fmt.Println(fmt.Sprintf(`EventFieldsEndIndex %v`, eventRecord.EventFieldsEndIndex)) */

	// Decode Topics
	if err := decoder.Decode(&eventRecord.Topics); err != nil {
		return EventRecord{}, newError(err, ErrorCode004)
	}
	/* 	fmt.Println("After topics") */

	eventRecord.EndIdx = uint32(decoder.Offset())
	eventRecord.AllBytes = decoder.ScaleBytes.Data
	eventRecord.Metadata = metadata
	eventRecord.Position = position

	// TODO
	/* 	fmt.Println(fmt.Sprintf(`Decoded %v, %v`, eventRecord.PalletName, eventRecord.EventName))
	   	fmt.Println(fmt.Sprintf(`Decoded %v`, eventRecord.Phase.ToString())) */
	/* 	fmt.Println(fmt.Sprintf(`EventFieldsStartIndexed %v`, eventRecord.EventFieldsStartIndex))
	   	fmt.Println(fmt.Sprintf(`EventFieldsEndIndex %v`, eventRecord.EventFieldsEndIndex)) */

	return eventRecord, nil
}

// Returns an array of E. Events that failed to decode are discarded.
// Use EventFindChecked if discarded events are necessary.
func EventFind[T interfaces.EventT](eventRecords EventRecords, target T) []T {
	var t T
	var result = []T{}
	for i := range eventRecords {
		if eventRecords[i].PalletIndex != target.PalletIndex() {
			continue
		}
		if eventRecords[i].EventIndex != target.EventIndex() {
			continue
		}

		var decoder = prim.NewDecoder(eventRecords[i].AllBytes[eventRecords[i].EventFieldsStartIndex:eventRecords[i].EventFieldsEndIndex], 0)
		if err := decoder.Decode(&t); err != nil {
			continue
		}
		result = append(result, t)
	}

	return result
}

// Returns an array of E.
// Some([]E) means we were able to decode all E events
// None means we failed to decode some E events.
func EventFindChecked[E interfaces.EventT](eventRecords EventRecords, target E) prim.Option[[]E] {
	var t E
	var result = []E{}
	for i := range eventRecords {
		if eventRecords[i].PalletIndex != target.PalletIndex() {
			continue
		}
		if eventRecords[i].EventIndex != target.EventIndex() {
			continue
		}

		var decoder = prim.NewDecoder(eventRecords[i].AllBytes[eventRecords[i].EventFieldsStartIndex:eventRecords[i].EventFieldsEndIndex], 0)
		if err := decoder.Decode(&t); err != nil {
			return prim.NewNone[[]E]()
		}

		result = append(result, t)
	}

	return prim.NewSome(result)
}

// Return None if the event has not been found.
// Returns Some(None) if the event has been found but we failed to decode it.
// Returns Some(E) if the event has been found and we decoded it.
func EventFindFirst[E interfaces.EventT](eventRecords EventRecords, target E) prim.Option[prim.Option[E]] {
	for i := range eventRecords {
		if eventRecords[i].PalletIndex != target.PalletIndex() {
			continue
		}
		if eventRecords[i].EventIndex != target.EventIndex() {
			continue
		}

		var t E
		var decoder = prim.NewDecoder(eventRecords[i].AllBytes[eventRecords[i].EventFieldsStartIndex:eventRecords[i].EventFieldsEndIndex], 0)
		if err := decoder.Decode(&t); err != nil {
			return prim.NewSome(prim.NewNone[E]())
		}
		return prim.NewSome(prim.NewSome(t))
	}

	return prim.NewNone[prim.Option[E]]()
}

// Return None if the event has not been found.
// Returns Some(None) if the event has been found but we failed to decode it.
// Returns Some(e) if the event has been found and we decoded it.
func EventFindLast[E interfaces.EventT](eventRecords EventRecords, target E) prim.Option[prim.Option[E]] {
	for i := range slices.Backward(eventRecords) {
		if eventRecords[i].PalletIndex != target.PalletIndex() {
			continue
		}
		if eventRecords[i].EventIndex != target.EventIndex() {
			continue
		}

		var t E
		var decoder = prim.NewDecoder(eventRecords[i].AllBytes[eventRecords[i].EventFieldsStartIndex:eventRecords[i].EventFieldsEndIndex], 0)
		if err := decoder.Decode(&t); err != nil {
			return prim.NewSome(prim.NewNone[E]())
		}
		return prim.NewSome(prim.NewSome(t))
	}

	return prim.NewNone[prim.Option[E]]()
}

func EventFilterByTxIndex(eventRecords EventRecords, txIndex uint32) EventRecords {
	var result = EventRecords{}
	for i := range eventRecords {
		if eventRecords[i].Phase.ApplyExtrinsic.IsNone() {
			continue
		}
		var elemTxIndex = eventRecords[i].Phase.ApplyExtrinsic.Unwrap()
		if elemTxIndex != txIndex {
			continue
		}

		result = append(result, eventRecords[i])
	}

	return result
}

func EventFilterSystemEvents(eventRecords EventRecords) EventRecords {
	var result = EventRecords{}
	for i := range eventRecords {
		if eventRecords[i].Phase.ApplyExtrinsic.IsSome() {
			continue
		}

		result = append(result, eventRecords[i])
	}

	return result
}
