package metadata

import (
	"bytes"
	"errors"
	"fmt"

	primitives "github.com/nmvalera/avail-go-sdk/primitives"

	gsrpcScale "github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	gsrpcTypes "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/itering/scale.go/utiles/uint128"
)

type Metadata struct {
	Value gsrpcTypes.Metadata
}

func NewMetadata(rawMetadata string) (Metadata, error) {
	scaleMetadata := primitives.Hex.FromHex(rawMetadata)

	metadata := Metadata{}
	if err := gsrpcScale.NewDecoder(bytes.NewReader(scaleMetadata)).Decode(&metadata.Value); err != nil {
		return Metadata{}, err
	}

	return metadata, nil
}

func (this *Metadata) PalletCallName(palletIndex uint8, callIndex uint8) (string, string, error) {
	pallet := this.FindPalletMetadata(palletIndex)
	if pallet == nil {
		return "", "", errors.New("Metadata Failure. Failed to find Pallet")
	}
	v14 := this.Value.AsMetadataV14

	if !pallet.HasCalls {
		return "", "", errors.New("Metadata Failure. Pallet has no Calls")
	}

	callId := pallet.Calls.Type.Int64()
	if typ, ok := v14.EfficientLookup[callId]; ok {
		if len(typ.Def.Variant.Variants) > 0 {
			for _, vars := range typ.Def.Variant.Variants {
				if uint8(vars.Index) != callIndex {
					continue
				}
				return string(pallet.Name), string(vars.Name), nil
			}
		}
	}

	return "", "", errors.New(fmt.Sprintf(`Metadata Failure. Failed to find pallet and event names. Pallet Index: %v, Call Index: %v`, palletIndex, callIndex))
}

func (this *Metadata) PalletEventName(palletIndex uint8, eventIndex uint8) (string, string, error) {
	pallet := this.FindPalletMetadata(palletIndex)
	if pallet == nil {
		return "", "", errors.New("Metadata Failure. Failed to find Pallet")
	}
	v14 := this.Value.AsMetadataV14

	if !pallet.HasEvents {
		return "", "", errors.New("Metadata Failure. Pallet has no events")
	}

	callId := pallet.Events.Type.Int64()
	if typ, ok := v14.EfficientLookup[callId]; ok {
		if len(typ.Def.Variant.Variants) > 0 {
			for _, vars := range typ.Def.Variant.Variants {
				if uint8(vars.Index) != eventIndex {
					continue
				}
				return string(pallet.Name), string(vars.Name), nil
			}
		}
	}

	return "", "", errors.New(fmt.Sprintf(`Metadata Failure. Failed to find pallet and event names. Pallet Index: %v, Event Index: %v`, palletIndex, eventIndex))
}

func (this *Metadata) FindPalletMetadata(palletIndex uint8) *gsrpcTypes.PalletMetadataV14 {
	v14 := this.Value.AsMetadataV14
	for _, pallet := range v14.Pallets {
		if uint8(pallet.Index) == palletIndex {
			return &pallet
		}
	}

	return nil
}

func (this *Metadata) DecodeEvent(palletIndex uint8, eventIndex uint8, decoder *primitives.Decoder) error {
	v14 := this.Value.AsMetadataV14
	pallet := this.FindPalletMetadata(palletIndex)
	if pallet == nil {
		return errors.New("Metadata Failure. Failed to find Pallet")
	}

	if !pallet.HasEvents {
		return errors.New("Metadata Failure. Pallet has no events")
	}

	callId := pallet.Events.Type.Int64()
	if typ, ok := v14.EfficientLookup[callId]; ok {
		if len(typ.Def.Variant.Variants) == 0 {
			return nil
		}

		for _, vars := range typ.Def.Variant.Variants {
			if uint8(vars.Index) != uint8(eventIndex) {
				continue
			}
			for _, field := range vars.Fields {
				if typ, ok1 := v14.EfficientLookup[field.Type.Int64()]; ok1 {
					this.decodeMetadataValue(decoder, typ, false)
				}
			}
		}
	} else {
		return errors.New(fmt.Sprintf(`Metadata Failure. No Type was found for the following id: %v`, callId))
	}

	return nil
}

func (this *Metadata) GetTypeFromId(id int64) *gsrpcTypes.Si1Type {
	v14 := this.Value.AsMetadataV14

	if typ, ok1 := v14.EfficientLookup[id]; ok1 {
		return typ
	}

	return nil
}

func (this *Metadata) decodeMetadataValue(decoder *primitives.Decoder, value *gsrpcTypes.Si1Type, isCompact bool) error {
	v14 := this.Value.AsMetadataV14

	/* 	path := ""
	   	for _, str := range value.Path {
	   		path += string(str) + " "
	   	}
	   	if path != "" {
	   		println(path)
	   	} */

	if value.Def.IsPrimitive {
		return DecodePrimitive(decoder, &value.Def.Primitive, isCompact)
	}

	if value.Def.IsArray {
		arr := value.Def.Array
		for i := 0; i < int(arr.Len); i++ {
			callId := value.Def.Array.Type.Int64()
			if typ, ok := v14.EfficientLookup[callId]; ok {
				if err := this.decodeMetadataValue(decoder, typ, isCompact); err != nil {
					return err
				}
			} else {
				return errors.New(fmt.Sprintf(`Metadata failure. Failed to find type with id: %v`, callId))
			}
		}

		return nil
	}

	if value.Def.IsTuple {
		tuple := value.Def.Tuple
		for _, elem := range tuple {
			callId := elem.Int64()
			if typ, ok := v14.EfficientLookup[callId]; ok {
				this.decodeMetadataValue(decoder, typ, isCompact)
			} else {
				return errors.New(fmt.Sprintf(`Metadata failure. Failed to find type with id: %v`, callId))
			}
		}

		return nil
	}

	if value.Def.IsSequence {
		len := primitives.CompactU32{}
		if err := decoder.Decode(&len); err != nil {
			return err
		}

		seq := value.Def.Sequence
		callId := seq.Type.Int64()
		for i := 0; i < int(len.Value); i++ {
			if typ, ok := v14.EfficientLookup[callId]; ok {
				if err := this.decodeMetadataValue(decoder, typ, isCompact); err != nil {
					return err
				}
			} else {
				return errors.New(fmt.Sprintf(`Metadata failure. Failed to find type with id: %v`, callId))
			}
		}

		return nil
	}

	if value.Def.IsCompact {
		callId := value.Def.Compact.Type.Int64()
		if typ, ok := v14.EfficientLookup[callId]; ok {
			if err := this.decodeMetadataValue(decoder, typ, true); err != nil {
				return err
			}
		} else {
			return errors.New(fmt.Sprintf(`Metadata failure. Failed to find type with id: %v`, callId))
		}

		return nil
	}

	if value.Def.IsComposite {
		com := value.Def.Composite
		for _, field := range com.Fields {
			callId := field.Type.Int64()
			if typ, ok := v14.EfficientLookup[callId]; ok {
				if err := this.decodeMetadataValue(decoder, typ, isCompact); err != nil {
					return err
				}
			} else {
				return errors.New(fmt.Sprintf(`Metadata failure. Failed to find type with id: %v`, callId))
			}
		}

		return nil
	}

	if value.Def.IsVariant {
		defVariant := value.Def.Variant
		index := uint8(0)
		if err := decoder.Decode(&index); err != nil {
			return err
		}

		found := false
		for _, variant := range defVariant.Variants {
			if uint8(variant.Index) != index {
				continue
			}
			found = true

			for _, field := range variant.Fields {
				callId := field.Type.Int64()
				if typ, ok := v14.EfficientLookup[callId]; ok {
					if err := this.decodeMetadataValue(decoder, typ, isCompact); err != nil {
						return err
					}
				} else {
					return errors.New(fmt.Sprintf(`Metadata failure. Failed to find type with id: %v`, callId))
				}
			}
		}

		if !found {
			return errors.New(fmt.Sprintf(`Metadata failure. Failed to find variant id %v`, index))
		}

		return nil
	}

	return errors.New(fmt.Sprintf(`Metadata failure. Don't know to how decode this type.`))
}

func DecodePrimitive(decoder *primitives.Decoder, value *gsrpcTypes.Si1TypeDefPrimitive, isCompact bool) error {
	if !isCompact {
		if int(value.Si0TypeDefPrimitive) == gsrpcTypes.IsBool {
			res := false
			if err := decoder.Decode(&res); err != nil {
				return err
			}
			return nil
		}

		if int(value.Si0TypeDefPrimitive) == gsrpcTypes.IsU8 {
			res := uint8(0)
			if err := decoder.Decode(&res); err != nil {
				return err
			}
			return nil
		}

		if int(value.Si0TypeDefPrimitive) == gsrpcTypes.IsU16 {
			res := uint16(0)
			if err := decoder.Decode(&res); err != nil {
				return err
			}
			return nil
		}

		if int(value.Si0TypeDefPrimitive) == gsrpcTypes.IsU32 {
			res := uint32(0)
			if err := decoder.Decode(&res); err != nil {
				return err
			}
			return nil
		}

		if int(value.Si0TypeDefPrimitive) == gsrpcTypes.IsU64 {
			res := uint64(0)
			if err := decoder.Decode(&res); err != nil {
				return err
			}
			return nil
		}

		if int(value.Si0TypeDefPrimitive) == gsrpcTypes.IsU128 {
			res := uint128.Uint128{}
			if err := decoder.Decode(&res); err != nil {
				return err
			}
			return nil
		}

		if int(value.Si0TypeDefPrimitive) == gsrpcTypes.IsStr {
			res := ""
			if err := decoder.Decode(&res); err != nil {
				return err
			}
			return nil
		}
	} else {
		if int(value.Si0TypeDefPrimitive) == gsrpcTypes.IsU32 {
			res := primitives.CompactU32{}
			if err := decoder.Decode(&res); err != nil {
				return err
			}
			return nil
		}

		if int(value.Si0TypeDefPrimitive) == gsrpcTypes.IsU64 {
			res := primitives.CompactU64{}
			if err := decoder.Decode(&res); err != nil {
				return err
			}
			return nil
		}

		if int(value.Si0TypeDefPrimitive) == gsrpcTypes.IsU128 {
			res := primitives.CompactU128{}
			if err := decoder.Decode(&res); err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New(fmt.Sprintf(`Metadata failure. Unknown primitive type: %v`, value.Si0TypeDefPrimitive))
}
