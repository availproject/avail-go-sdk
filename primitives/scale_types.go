package primitives

import (
	"fmt"

	"github.com/itering/scale.go/utiles/uint128"
)

type Option[T any] struct {
	value T
	isSet bool
}

func (o *Option[T]) EncodeTo(dest *string) {
	if !o.isSet {
		Encoder.EncodeTo(uint8(0), dest)
		return
	}
	Encoder.EncodeTo(uint8(1), dest)
	Encoder.EncodeTo(o.value, dest)
}

func (o *Option[T]) Decode(decoder *Decoder) error {
	hasValue := uint8(0)
	if err := decoder.Decode(&hasValue); err != nil {
		return err
	}
	if hasValue == 1 {
		o.isSet = true
		if err := decoder.Decode(&o.value); err != nil {
			return err
		}
	} else {
		o.isSet = false
	}

	return nil
}

func (o *Option[T]) Set(value T) {
	o.value = value
	o.isSet = true
}

func (o *Option[T]) Unset() {
	var t T
	o.value = t
	o.isSet = false
}

func (o Option[T]) IsSome() bool {
	return o.isSet
}

func (o Option[T]) IsNone() bool {
	return !o.isSet
}

func (o Option[T]) String() string {
	if o.isSet {
		return fmt.Sprintf("Some(%v)", o.value)
	} else {
		return "None"
	}
}

func (o Option[T]) ToString() string {
	return o.String()
}

func (o Option[T]) ToHuman() string {
	return o.String()
}

// This function does not panic when no value is set.
//
// If Set, returns set value.
// If Not Set, returns default value.
func (o Option[T]) Unwrap() T {
	if o.isSet == false {
		var t T
		return t
	}
	return o.value
}

// This function will panic when no value is set.
//
// If Set, returns set value.
// If Not Set, panics
func (o Option[T]) UnsafeUnwrap() T {
	if o.isSet == false {
		panic("Option is not set.")
	}
	return o.value
}

// This function does not panic when no value is set.
//
// If Set, returns set value.
// If Not Set, returns value specified as parameter.
func (o Option[T]) UnwrapOr(elseValue T) T {
	if o.isSet == false {
		return elseValue
	}
	return o.value
}

func Some[T any](value T) Option[T] {
	option := Option[T]{}
	option.Set(value)
	return option
}

func None[T any]() Option[T] {
	return Option[T]{}
}

type CompactU32 struct {
	Value uint32 `scale:"compact"`
}

type CompactU64 struct {
	Value uint64 `scale:"compact"`
}

type CompactU128 struct {
	Value uint128.Uint128 `scale:"compact"`
}
