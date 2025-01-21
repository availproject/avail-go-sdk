package sdk

import (
	"errors"
	"fmt"

	interfaces "github.com/nmvalera/avail-go-sdk/interfaces"
	meta "github.com/nmvalera/avail-go-sdk/metadata"
	prim "github.com/nmvalera/avail-go-sdk/primitives"
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
	decoder.Decode(&compactEventCount)

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
		return EventPhase{}, err
	}

	switch eventPhase.VariantIndex {
	case 0:
		var value = uint32(0)
		if err := decoder.Decode(&value); err != nil {
			return EventPhase{}, err
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

func NewEventRecord(decoder *prim.Decoder, position uint32, metadata *meta.Metadata) (EventRecord, error) {
	var eventRecord = EventRecord{}

	eventRecord.StartIdx = uint32(decoder.Offset())

	/* 	println("Before Event Phase") */
	eventPhase, err := DecodeEventPhase(decoder)
	if err != nil {
		return EventRecord{}, err
	}
	eventRecord.Phase = eventPhase
	/* 	println("Before After Phase") */

	// Pallet and Event Index Decoding

	eventRecord.EventStartIndex = uint32(decoder.Offset())
	if err := decoder.Decode(&eventRecord.PalletIndex); err != nil {
		return EventRecord{}, err
	}
	if err := decoder.Decode(&eventRecord.EventIndex); err != nil {
		return EventRecord{}, err
	}

	// Decoding Pallet and Event Names
	palletName, eventName, err := metadata.PalletEventName(eventRecord.PalletIndex, eventRecord.EventIndex)
	if err != nil {
		return EventRecord{}, err
	}
	eventRecord.PalletName = palletName
	eventRecord.EventName = eventName

	// println(fmt.Sprintf(`Decoding %v, %v %v`, eventRecord.PalletName, eventRecord.EventName, eventRecord.Phase.ToString()))

	// Decode Fields
	eventRecord.EventFieldsStartIndex = uint32(decoder.Offset())
	if err := metadata.DecodeEvent(eventRecord.PalletIndex, eventRecord.EventIndex, decoder); err != nil {
		return EventRecord{}, err
	}
	eventRecord.EventFieldsEndIndex = uint32(decoder.Offset())
	/* 	println("Before topics")

	   	println(fmt.Sprintf(`EventFieldsStartIndexed %v`, eventRecord.EventFieldsStartIndex))
	   	println(fmt.Sprintf(`EventFieldsEndIndex %v`, eventRecord.EventFieldsEndIndex)) */

	// Decode Topics
	if err := decoder.Decode(&eventRecord.Topics); err != nil {
		return EventRecord{}, err
	}
	/* 	println("After topics") */

	eventRecord.EndIdx = uint32(decoder.Offset())
	eventRecord.AllBytes = decoder.ScaleBytes.Data
	eventRecord.Metadata = metadata
	eventRecord.Position = position

	// TODO
	/* 	println(fmt.Sprintf(`Decoded %v, %v`, eventRecord.PalletName, eventRecord.EventName))
	   	println(fmt.Sprintf(`Decoded %v`, eventRecord.Phase.ToString())) */
	/* 	println(fmt.Sprintf(`EventFieldsStartIndexed %v`, eventRecord.EventFieldsStartIndex))
	   	println(fmt.Sprintf(`EventFieldsEndIndex %v`, eventRecord.EventFieldsEndIndex)) */

	return eventRecord, nil
}

func EventFindFirst[T interfaces.EventT](eventRecords EventRecords, target T) prim.Option[T] {
	event, err := EventFindFirstChecked(eventRecords, target)
	if err != nil {
		return prim.NewNone[T]()
	}
	return event
}

func EventFindFirstChecked[T interfaces.EventT](eventRecords EventRecords, target T) (prim.Option[T], error) {
	for _, elem := range eventRecords {
		if elem.PalletIndex != target.PalletIndex() {
			continue
		}
		if elem.EventIndex != target.EventIndex() {
			continue
		}

		var t T
		var decoder = prim.NewDecoder(elem.AllBytes[elem.EventFieldsStartIndex:elem.EventFieldsEndIndex], 0)
		if err := decoder.Decode(&t); err != nil {
			return prim.NewNone[T](), err
		}

		return prim.NewSome(t), nil
	}

	return prim.NewNone[T](), nil
}

func EventFindAll[T interfaces.EventT](eventRecords EventRecords, target T) []T {
	events, err := EventFindAllChecked(eventRecords, target)
	if err != nil {
		return []T{}
	}
	return events
}

func EventFindAllChecked[T interfaces.EventT](eventRecords EventRecords, target T) ([]T, error) {
	var t T
	var result = []T{}
	for _, elem := range eventRecords {
		if elem.PalletIndex != target.PalletIndex() {
			continue
		}
		if elem.EventIndex != target.EventIndex() {
			continue
		}

		var decoder = prim.NewDecoder(elem.AllBytes[elem.EventFieldsStartIndex:elem.EventFieldsEndIndex], 0)
		if err := decoder.Decode(&t); err != nil {
			return []T{}, err
		}
		result = append(result, t)
	}

	return result, nil
}

func EventFindLast[T interfaces.EventT](eventRecords EventRecords, target T) prim.Option[T] {
	event, err := EventFindLastChecked(eventRecords, target)
	if err != nil {
		return prim.NewNone[T]()
	}
	return event
}

func EventFindLastChecked[T interfaces.EventT](eventRecords EventRecords, target T) (prim.Option[T], error) {
	result, err := EventFindAllChecked(eventRecords, target)
	if err != nil {
		return prim.NewNone[T](), err
	}
	if len(result) == 0 {
		return prim.NewNone[T](), nil
	}

	return prim.NewSome(result[len(result)-1]), nil
}

func FilterByTxIndex(eventRecords EventRecords, txIndex uint32) EventRecords {
	var result = EventRecords{}
	for _, elem := range eventRecords {
		if elem.Phase.ApplyExtrinsic.IsNone() {
			continue
		}
		var elemTxIndex = elem.Phase.ApplyExtrinsic.Unwrap()
		if elemTxIndex != txIndex {
			continue
		}

		result = append(result, elem)
	}

	return result
}

func FilterSystemEvents(eventRecords EventRecords, txIndex uint32) EventRecords {
	var result = EventRecords{}
	for _, elem := range eventRecords {
		if elem.Phase.ApplyExtrinsic.IsSome() {
			continue
		}

		result = append(result, elem)
	}

	return result
}
