package primitives

import (
	"fmt"
	"math"
	"math/bits"

	"github.com/itering/scale.go/utiles"
	"github.com/itering/scale.go/utiles/uint128"
	"github.com/vedhavyas/go-subkey/v2"
	"golang.org/x/crypto/blake2b"
)

// Do not change the order of field members.
type Additional struct {
	SpecVersion uint32
	TxVersion   uint32
	GenesisHash H256
	ForkHash    H256
}

// Do not change the order of field members.
type Call struct {
	PalletIndex uint8
	CallIndex   uint8
	Fields      AlreadyEncoded
}

func NewCall(palletIndex uint8, callIndex uint8, fields AlreadyEncoded) Call {
	return Call{
		PalletIndex: palletIndex,
		CallIndex:   callIndex,
		Fields:      fields,
	}
}

func (this *Call) Decode(decoder *Decoder) error {
	// Call Index
	if err := decoder.Decode(&this.PalletIndex); err != nil {
		return err
	}
	if err := decoder.Decode(&this.CallIndex); err != nil {
		return err
	}

	// Call Data
	dataBytes := decoder.NextBytes(decoder.RemainingLength())
	this.Fields = AlreadyEncoded{Value: Hex.ToHex(dataBytes)}
	return nil
}

// Do not change the order of field members.
type Era struct {
	IsImmortal bool
	Period     uint64
	Phase      uint64
}

func (this Era) ToHuman() string {
	return this.String()
}

func (this Era) ToString() string {
	return this.String()
}
func (this Era) String() string {
	if this.IsImmortal {
		return fmt.Sprintf("Immortal")
	}

	return fmt.Sprintf("Mortal: {period: %v, phase: %v}", this.Period, this.Phase)
}

func (this *Era) EncodeTo(dest *string) {
	if this.IsImmortal {
		Encoder.EncodeTo(uint8(0), dest)
		return
	}

	quantizeFactor := math.Max(float64(this.Period>>12), 1)
	trailingZeros := bits.TrailingZeros16(uint16(this.Period))
	encoded := uint16(float64(this.Phase)/quantizeFactor)<<4 | uint16(math.Min(15, math.Max(1, float64(trailingZeros-1))))

	first := byte(encoded & 0xff)
	second := byte(encoded >> 8)

	Encoder.EncodeTo(uint8(first), dest)
	Encoder.EncodeTo(uint8(second), dest)
}

func (this *Era) Decode(decoder *Decoder) error {
	*this = Era{}

	first := uint8(0)
	if err := decoder.Decode(&first); err != nil {
		return err
	}

	if first == 0 {
		this.IsImmortal = true
		return nil
	}

	second := uint8(0)
	if err := decoder.Decode(&second); err != nil {
		return err
	}

	encoded := uint16(first) | (uint16(second) << 8)

	trailingZeros := uint16(encoded&0xF) + 1 // Lower 4 bits + 1
	quantizedPhase := encoded >> 4           // Upper 12 bits

	quantizeFactor := math.Max(float64(this.Period>>12), 1)
	this.Phase = uint64(float64(quantizedPhase) * quantizeFactor)
	this.Period = uint64(uint16(1 << trailingZeros))

	return nil
}

// Mortal describes a mortal era based on a period of validity and a block number on which it should start
func NewEra(period uint64, blockNumber uint64) Era {
	calPeriod := uint64(math.Pow(2, math.Ceil(math.Log2(float64(period)))))
	if calPeriod < 4 {
		calPeriod = 4
	}
	if calPeriod > (1 << 16) {
		calPeriod = 1 << 16
	}

	phase := blockNumber % calPeriod
	quantize_factor := math.Max(float64(calPeriod>>12), 1)
	quantized_phase := float64(phase) / quantize_factor * quantize_factor

	return Era{
		Period: calPeriod,
		Phase:  uint64(quantized_phase),
	}
}

// Do not change the order of field members.
type Extra struct {
	Era   Era
	Nonce uint32          `scale:"compact"`
	Tip   uint128.Uint128 `scale:"compact"`
	AppId uint32          `scale:"compact"`
}

type UnsignedPayload struct {
	Call       Call
	Extra      Extra
	Additional Additional
}

func (this UnsignedPayload) Encode() UnsignedEncodedPayload {
	return UnsignedEncodedPayload{
		Call:       AlreadyEncoded{Value: Encoder.Encode(this.Call)},
		Extra:      AlreadyEncoded{Value: Encoder.Encode(this.Extra)},
		Additional: AlreadyEncoded{Value: Encoder.Encode(this.Additional)},
	}
}

type UnsignedEncodedPayload struct {
	Call       AlreadyEncoded
	Extra      AlreadyEncoded
	Additional AlreadyEncoded
}

func (this *UnsignedEncodedPayload) Sign(signer subkey.KeyPair) ([]byte, error) {
	data := ""
	Encoder.EncodeTo(this.Call, &data)
	Encoder.EncodeTo(this.Extra, &data)
	Encoder.EncodeTo(this.Additional, &data)

	decodedData := utiles.HexToBytes(data)

	if len(decodedData) > 256 {
		blakeSum := blake2b.Sum256(decodedData)
		return signer.Sign([]byte(blakeSum[:]))
	} else {
		return signer.Sign(decodedData)
	}

}
