package primitives

import (
	"fmt"

	"github.com/itering/scale.go/utiles/uint128"
)

type Option[T any] struct {
	value T
	isSet bool
}

func (this *Option[T]) EncodeTo(dest *string) {
	if !this.isSet {
		Encoder.EncodeTo(uint8(0), dest)
		return
	}
	Encoder.EncodeTo(uint8(1), dest)
	Encoder.EncodeTo(this.value, dest)
}

func (this *Option[T]) Decode(decoder *Decoder) error {
	hasValue := uint8(0)
	if err := decoder.Decode(&hasValue); err != nil {
		return err
	}
	if hasValue == 1 {
		this.isSet = true
		if err := decoder.Decode(&this.value); err != nil {
			return err
		}
	} else {
		this.isSet = false
	}

	return nil
}

func (this *Option[T]) Set(value T) {
	this.value = value
	this.isSet = true
}

func (this *Option[T]) Unset() {
	var t T
	this.value = t
	this.isSet = false
}

func (this Option[T]) IsSome() bool {
	return this.isSet
}

func (this Option[T]) IsNone() bool {
	return !this.isSet
}

func (this Option[T]) String() string {
	if this.isSet {
		return fmt.Sprintf("%v", this.value)
	} else {
		return "None"
	}
}

func (this Option[T]) ToString() string {
	return this.String()
}

func (this Option[T]) ToHuman() string {
	return this.String()
}

// This function does not panic when no value is set.
//
// If Set, returns set value.
// If Not Set, returns default value.
func (this Option[T]) Unwrap() T {
	if this.isSet == false {
		var t T
		return t
	}
	return this.value
}

// This function will panic when no value is set.
//
// If Set, returns set value.
// If Not Set, panics
func (this Option[T]) UnsafeUnwrap() T {
	if this.isSet == false {
		panic("Option is not set.")
	}
	return this.value
}

// This function does not panic when no value is set.
//
// If Set, returns set value.
// If Not Set, returns value specified as parameter.
func (this Option[T]) UnwrapOr(elseValue T) T {
	if this.isSet == false {
		return elseValue
	}
	return this.value
}

func NewSome[T any](value T) Option[T] {
	option := Option[T]{}
	option.Set(value)
	return option
}

func NewNone[T any]() Option[T] {
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
