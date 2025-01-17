package primitives

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/xxhash"
	"golang.org/x/crypto/blake2b"
)

func TwoX128(input string) []byte {
	return xxhash.New128([]byte(input)).Sum(nil)
}

// blake2_128_concat computes the Blake2_128Concat hash
func Blake2_128_Concat(input []byte) []byte {
	// Create a Blake2b hasher with a 16-byte output (128-bit).
	hasher, err := blake2b.New(16, nil)
	if err != nil {
		panic(err) // Ensure the hasher initializes correctly
	}

	// Write the input to the hasher and compute the hash
	hasher.Write(input)
	hash := hasher.Sum(nil) // This produces a 16-byte (128-bit) hash

	// Concatenate the hash with the original input bytes
	return append(hash, input...)
}

func DecodeBlake2_128Concat(data []byte) []byte {
	// Blake2_128Concat keys are in the format:
	// [16-byte Blake2_128 hash | original key bytes]
	if len(data) <= 16 {
		panic("Invalid Blake2_128Concat key format")
	}
	return data[16:] // Return the original key bytes
}
