# efm-go

A tiny Go library that encodes and decodes data using an EFM (Eight-to-Fourteen) style mapping.

- Input is processed in fixed-size blocks.
  - Encode: 8 input bytes -> 14 output bytes
  - Decode: 14 input bytes -> 8 output bytes
- Zero-allocation hot paths avoided where correctness is prioritized; maps are precomputed for fast lookups.

## Installation

This project uses Go modules. In your project, add it as a dependency. If you use this repository as-is (local module name `efm-go`), you can import it with an alias:

```go
import (
    efm_go "efm-go"
)
```

If you fork or publish under a VCS host (e.g., GitHub), replace `efm-go` above with your module path.

## Quick Start

```go
package main

import (
    "fmt"
    efm_go "efm-go"
)

func main() {
    efm := efm_go.New()

    // Encode: input length MUST be a multiple of 8
    plain := []byte("12345678")
    encoded, err := efm.Encode(plain)
    if err != nil {
        panic(err)
    }
    fmt.Printf("encoded (%d bytes): %v\n", len(encoded), encoded)

    // Decode: input length MUST be a multiple of 14
    decoded, err := efm.Decode(encoded)
    if err != nil {
        panic(err)
    }
    fmt.Printf("decoded: %s\n", string(decoded))
}
```

## API

- New() *EFM
  - Constructs a new encoder/decoder instance with precomputed encode/decode tables.
- (EFM) Encode(data []byte) ([]byte, error)
  - Returns EFM-encoded bytes.
  - Returns an error if len(data) is not a multiple of 8.
- (EFM) Decode(data []byte) ([]byte, error)
  - Returns decoded bytes.
  - Returns an error if len(data) is not a multiple of 14 or the input contains invalid quattuordecuple segments.

### Data layout details

Internally, each 8-byte block is transformed into a 112-bit (14-byte) value using a look-up table. Encoding/decoding is performed quattuordecuple (14-bit) at a time, using a 128-bit accumulator to pack/unpack bits efficiently.

Constraints to remember:
- Encode input length must be divisible by 8.
- Decode input length must be divisible by 14.

## Examples

Encoding a single 8-byte block:

```go
plain := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
ef := efm_go.New()
encoded, err := ef.Encode(plain)
if err != nil { /* handle */ }
// len(encoded) == 14
```

Decoding back:

```go
decoded, err := ef.Decode(encoded)
if err != nil { /* handle */ }
// bytes.Equal(decoded, plain) == true
```

## Testing

Run the test suite:

```bash
go test ./...
```

## License

This project is licensed under the terms of the LICENSE.md file in this repository.

## Notes

- The package name is `efm_go` (underscore), so when importing via its module path, consider aliasing as `efm_go` for idiomatic usage.
- The library depends on `github.com/davidminor/uint128` for 128-bit operations.
