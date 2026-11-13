package primitives

import (
	"runtime"
	"testing"
)

// slice() decodes a SCALE compact length and then calls
// reflect.MakeSlice(type, n, n) up front, before reading any element. A crafted
// compact length (up to ~4.29e9) therefore triggers a huge allocation from only
// a few bytes of input. These tests pin that down and verify the fix.

// A compact length of max-uint32 with no element bytes: before the fix this
// makes a ~4.29e9-element slice up front -> "allocation size out of range"
// panic. After the fix it must return an error without panicking.
func TestSliceRejectsHugeCompactLength(t *testing.T) {
	blob := Hex.FromHex(Encoder.Encode(CompactU32{Value: 4294967295}))

	var panicked interface{}
	var err error
	func() {
		defer func() { panicked = recover() }()
		dec := NewDecoder(blob, 0)
		var out []uint32
		err = dec.Decode(&out)
	}()

	t.Logf("len(blob)=%d  err=%v  panicked=%v", len(blob), err, panicked)
	if panicked != nil {
		t.Fatalf("decode panicked on a huge declared slice length (DoS): %v", panicked)
	}
	if err == nil {
		t.Fatalf("expected an error for a slice length beyond the input, got nil")
	}
}

// A compact length of 100,000,000 with no element bytes. Before the fix this
// pre-allocates a 100M-element []uint32 (~400 MB) before failing to read the
// elements. After the fix it must reject in O(1) without the large allocation.
func TestSliceDoesNotPreallocateBeyondRemaining(t *testing.T) {
	blob := Hex.FromHex(Encoder.Encode(CompactU32{Value: 100000000}))

	runtime.GC()
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)

	dec := NewDecoder(blob, 0)
	var out []uint32
	err := dec.Decode(&out)

	runtime.ReadMemStats(&after)
	alloc := int64(after.TotalAlloc - before.TotalAlloc)
	t.Logf("err=%v  bytes allocated during decode=%d", err, alloc)

	if err == nil {
		t.Fatalf("expected an error for a slice length beyond the input, got nil")
	}
	if alloc > 10*1024*1024 {
		t.Fatalf("pre-allocated %d bytes before validating the declared length (DoS)", alloc)
	}
}

// Honestly-encoded input must still round-trip unchanged after the fix.
func TestSliceStillDecodesValidInput(t *testing.T) {
	in := []uint32{1, 2, 3, 4, 5}
	blob := Hex.FromHex(Encoder.Encode(in))

	dec := NewDecoder(blob, 0)
	var out []uint32
	if err := dec.Decode(&out); err != nil {
		t.Fatalf("decode of valid input failed: %v", err)
	}
	if len(out) != len(in) {
		t.Fatalf("round-trip length mismatch: got %d want %d", len(out), len(in))
	}
	for i := range in {
		if out[i] != in[i] {
			t.Fatalf("round-trip value mismatch at %d: got %d want %d", i, out[i], in[i])
		}
	}
}
