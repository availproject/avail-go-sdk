package primitives

type AlreadyEncoded struct {
	Value string
}

func (this *AlreadyEncoded) EncodeTo(dest *string) {
	*dest += this.Value
}

func (this *AlreadyEncoded) ToHex() string {
	return this.Value
}

func (this *AlreadyEncoded) ToHexWith0x() string {
	return "0x" + this.Value
}

func (this *AlreadyEncoded) ToBytes() []byte {
	return Hex.FromHex(this.Value)
}
