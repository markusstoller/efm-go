package efm_go

import (
	"encoding/binary"
	"fmt"
	"github.com/davidminor/uint128"
)

const invalidCode = 0xFFFF
const encodeBufferSize = 256
const decodeBufferSize = 0x4000
const quattuordecupleSize = 14
const startPt = 112 - quattuordecupleSize
const bufferMask = 0x3FFF
const window = 8

type EFM struct {
	decodeMapping [decodeBufferSize]uint16
	encodeMapping [encodeBufferSize]uint16
}

func New() *EFM {
	encodeMapping := [encodeBufferSize]uint16{invalidCode}
	decodeMapping := [decodeBufferSize]uint16{invalidCode}

	efmRawTable := []uint16{
		0b01001000100000,
		0b10000100000000,
		0b10010000100000,
		0b10001000100000,
		0b01000100000000,
		0b00000100010000,
		0b00010000100000,
		0b00100100000000,
		0b01001001000000,
		0b10000001000000,
		0b10010001000000,
		0b10001001000000,
		0b01000001000000,
		0b00000001000000,
		0b00010001000000,
		0b00100001000000,
		0b10000000100000,
		0b10000010000000,
		0b10010010000000,
		0b00100000100000,
		0b01000010000000,
		0b00000010000000,
		0b00010010000000,
		0b00100010000000,
		0b01001000010000,
		0b10000000010000,
		0b10010000010000,
		0b10001000010000,
		0b01000000010000,
		0b00001000010000,
		0b00010000010000,
		0b00100000010000,
		0b00000000100000,
		0b10000100001000,
		0b00001000100000,
		0b00100100100000,
		0b01000100001000,
		0b00000100001000,
		0b01000000100000,
		0b00100100001000,
		0b01001001001000,
		0b10000001001000,
		0b10010001001000,
		0b10001001001000,
		0b01000001001000,
		0b00000001001000,
		0b00010001001000,
		0b00100001001000,
		0b00000100000000,
		0b10000010001000,
		0b10010010001000,
		0b10000100010000,
		0b01000010001000,
		0b00000010001000,
		0b00010010001000,
		0b00100010001000,
		0b01001000001000,
		0b10000000001000,
		0b10010000001000,
		0b10001000001000,
		0b01000000001000,
		0b00001000001000,
		0b00010000001000,
		0b00100000001000,
		0b01001000100100,
		0b10000100100100,
		0b10010000100100,
		0b10001000100100,
		0b01000100100100,
		0b00000000100100,
		0b00010000100100,
		0b00100100100100,
		0b01001001000100,
		0b10000001000100,
		0b10010001000100,
		0b10001001000100,
		0b01000001000100,
		0b00000001000100,
		0b00010001000100,
		0b00100001000100,
		0b10000000100100,
		0b10000010000100,
		0b10010010000100,
		0b00100000100100,
		0b01000010000100,
		0b00000010000100,
		0b00010010000100,
		0b00100010000100,
		0b01001000000100,
		0b10000000000100,
		0b10010000000100,
		0b10001000000100,
		0b01000000000100,
		0b00001000000100,
		0b00010000000100,
		0b00100000000100,
		0b01001000100010,
		0b10000100100010,
		0b10010000100010,
		0b10001000100010,
		0b01000100100010,
		0b00000000100010,
		0b01000000100100,
		0b00100100100010,
		0b01001001000010,
		0b10000001000010,
		0b10010001000010,
		0b10001001000010,
		0b01000001000010,
		0b00000001000010,
		0b00010001000010,
		0b00100001000010,
		0b10000000100010,
		0b10000010000010,
		0b10010010000010,
		0b00100000100010,
		0b01000010000010,
		0b00000010000010,
		0b00010010000010,
		0b00100010000010,
		0b01001000000010,
		0b00001001001000,
		0b10010000000010,
		0b10001000000010,
		0b01000000000010,
		0b00001000000010,
		0b00010000000010,
		0b00100000000010,
		0b01001000100001,
		0b10000100100001,
		0b10010000100001,
		0b10001000100001,
		0b01000100100001,
		0b00000000100001,
		0b00010000100001,
		0b00100100100001,
		0b01001001000001,
		0b10000001000001,
		0b10010001000001,
		0b10001001000001,
		0b01000001000001,
		0b00000001000001,
		0b00010001000001,
		0b00100001000001,
		0b10000000100001,
		0b10000010000001,
		0b10010010000001,
		0b00100000100001,
		0b01000010000001,
		0b00000010000001,
		0b00010010000001,
		0b00100010000001,
		0b01001000000001,
		0b10000010010000,
		0b10010000000001,
		0b10001000000001,
		0b01000010010000,
		0b00001000000001,
		0b00010000000001,
		0b00100010010000,
		0b00001000100001,
		0b10000100001001,
		0b01000100010000,
		0b00000100100001,
		0b01000100001001,
		0b00000100001001,
		0b01000000100001,
		0b00100100001001,
		0b01001001001001,
		0b10000001001001,
		0b10010001001001,
		0b10001001001001,
		0b01000001001001,
		0b00000001001001,
		0b00010001001001,
		0b00100001001001,
		0b00000100100000,
		0b10000010001001,
		0b10010010001001,
		0b00100100010000,
		0b01000010001001,
		0b00000010001001,
		0b00010010001001,
		0b00100010001001,
		0b01001000001001,
		0b10000000001001,
		0b10010000001001,
		0b10001000001001,
		0b01000000001001,
		0b00001000001001,
		0b00010000001001,
		0b00100000001001,
		0b01000100100000,
		0b10000100010001,
		0b10010010010000,
		0b00001000100100,
		0b01000100010001,
		0b00000100010001,
		0b00010010010000,
		0b00100100010001,
		0b00001001000001,
		0b10000100000001,
		0b00001001000100,
		0b00001001000000,
		0b01000100000001,
		0b00000100000001,
		0b00000010010000,
		0b00100100000001,
		0b00000100100100,
		0b10000010010001,
		0b10010010010001,
		0b10000100100000,
		0b01000010010001,
		0b00000010010001,
		0b00010010010001,
		0b00100010010001,
		0b01001000010001,
		0b10000000010001,
		0b10010000010001,
		0b10001000010001,
		0b01000000010001,
		0b00001000010001,
		0b00010000010001,
		0b00100000010001,
		0b01000100000010,
		0b00000100000010,
		0b10000100010010,
		0b00100100000010,
		0b01000100010010,
		0b00000100010010,
		0b01000000100010,
		0b00100100010010,
		0b10000100000010,
		0b10000100000100,
		0b00001001001001,
		0b00001001000010,
		0b01000100000100,
		0b00000100000100,
		0b00010000100010,
		0b00100100000100,
		0b00000100100010,
		0b10000010010010,
		0b10010010010010,
		0b00001000100010,
		0b01000010010010,
		0b00000010010010,
		0b00010010010010,
		0b00100010010010,
		0b01001000010010,
		0b10000000010010,
		0b10010000010010,
		0b10001000010010,
		0b01000000010010,
		0b00001000010010,
		0b00010000010010,
		0b00100000010010,
	}

	for i := 0; i < len(efmRawTable); i++ {
		decodeMapping[efmRawTable[i]] = uint16(i)
		encodeMapping[i] = efmRawTable[i]
	}
	return &EFM{
		encodeMapping: encodeMapping,
		decodeMapping: decodeMapping,
	}
}

// encodeQuattuordecuple converts an 8-byte input slice into a uint128 value by encoding each byte and shifting into position.
// Returns an error if the input slice length is less than 8.
func (e *EFM) encodeQuattuordecuple(data []byte) (uint128.Uint128, error) {
	result := uint128.Uint128{}
	if len(data) < window {
		return result, fmt.Errorf("data length must be at least 8")
	}

	for i := 0; i < window; i++ {
		shiftAmount := startPt - i*quattuordecupleSize
		// Convert val to uint128 and shift it to the correct position
		valAs128 := uint128.Uint128{L: uint64(e.encodeMapping[data[i]])}
		shiftedVal := valAs128.ShiftLeft(uint(shiftAmount))

		// Combine
		result = result.Or(shiftedVal)
	}
	return result, nil
}

// Encode encodes the provided data using EFM scheme, requiring the input length to be a multiple of 8 bytes.
// It returns the encoded byte slice or an error if the input length is invalid.
func (e *EFM) Encode(data []byte) ([]byte, error) {
	if len(data)%window != 0 {
		return nil, fmt.Errorf("data length must be multiple of 8")
	}

	// Pre-allocate result slice with exact capacity needed
	// Each 8-byte chunk produces 14 bytes of output (6 + 8)
	numChunks := len(data) / window
	result := make([]byte, 0, numChunks*quattuordecupleSize)

	// Reuse buffer for binary encoding to avoid repeated allocations
	buf := make([]byte, window)

	for i := 0; i < len(data); i += window {
		encoded, _ := e.encodeQuattuordecuple(data[i : i+window])

		// Encode high part (only need 6 bytes from position 2)
		binary.BigEndian.PutUint64(buf, encoded.H)
		result = append(result, buf[2:]...)

		// Encode low part (all 8 bytes)
		binary.BigEndian.PutUint64(buf, encoded.L)
		result = append(result, buf...)
	}
	return result, nil
}

// decodeQuattuordecuple converts a uint128 encoded value into a slice of 8 decoded bytes using EFM decoding rules.
// Returns an error if any segment of the encoded value is invalid.
func (e *EFM) decodeQuattuordecuple(encodedValue uint128.Uint128) ([]byte, error) {
	result := make([]byte, window)

	for i := 0; i < window; i++ {
		shiftAmount := startPt - i*quattuordecupleSize
		decoded := e.decodeMapping[encodedValue.ShiftRight(uint(shiftAmount)).L&bufferMask]

		if decoded == invalidCode {
			return nil, fmt.Errorf("invalid value")
		}

		result[i] = byte(decoded)
	}

	return result, nil
}

// Decode decodes the given EFM-encoded byte slice, ensuring the input length is a multiple of 14 bytes. Returns decoded bytes or an error.
func (e *EFM) Decode(data []byte) ([]byte, error) {
	result := make([]byte, 0, len(data)/quattuordecupleSize)

	if len(data)%quattuordecupleSize != 0 {
		return nil, fmt.Errorf("data length must be multiple of %d", quattuordecupleSize)
	}

	for i := 0; i < len(data); i += quattuordecupleSize {
		uint128Data := uint128.Uint128{H: binary.BigEndian.Uint64(data[i+0:i+window]) >> 16, L: binary.BigEndian.Uint64(data[i+6 : i+14])}
		decoded, err := e.decodeQuattuordecuple(uint128Data)
		if err != nil {
			return nil, err
		}
		result = append(result, decoded...)

	}

	return result, nil
}
