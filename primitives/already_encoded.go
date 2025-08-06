package primitives

type AlreadyEncoded struct {
	Value string
}

func (a *AlreadyEncoded) EncodeTo(dest *string) {
	*dest += a.Value
}

func (a *AlreadyEncoded) ToHex() string {
	return a.Value
}

func (a *AlreadyEncoded) ToHexWith0x() string {
	return "0x" + a.Value
}

func (a *AlreadyEncoded) ToBytes() []byte {
	return Hex.FromHex(a.Value)
}
