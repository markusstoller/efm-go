package efm_go

import (
	"github.com/davidminor/uint128"
	"testing"
)

func TestEFM_decodeQuattuordecuple(t *testing.T) {
	tests := []struct {
		name      string
		input     uint128.Uint128
		want      []byte
		expectErr bool
	}{
		{
			name:      "valid decoding",
			input:     uint128.Uint128{H: 145137984342281, L: 2486275223393945860},
			want:      []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
			expectErr: false,
		},
		{
			name:      "invalid encoded value",
			input:     uint128.Uint128{H: 0, L: 0xFFFFFFFFFFFFFFFF},
			want:      nil,
			expectErr: true,
		},
		{
			name:      "zero encoded value",
			input:     uint128.Uint128{H: 0, L: 0},
			want:      []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()
			got, err := e.decodeQuattuordecuple(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("EFM.decodeQuattuordecuple() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !tt.expectErr && !compareSlices(got, tt.want) {
				t.Errorf("EFM.decodeQuattuordecuple() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestEFM_Decode(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    []byte
		wantErr bool
	}{
		{
			name: "valid data decoding",

			input:   []byte{130, 34, 72, 136, 68, 16, 136, 2, 32, 72, 130, 34, 18, 8},
			want:    []byte("12345678"),
			wantErr: false,
		},
		{
			name:    "input length not divisible by 14",
			input:   []byte{0x01, 0x23, 0x45},
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []byte{},
			want:    []byte{},
			wantErr: false,
		},
		{
			name:    "invalid input",
			input:   append([]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D}, []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D}...),
			want:    append([]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}...),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()
			got, err := e.Decode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("EFM.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !compareSlices(got, tt.want) {
				t.Errorf("EFM.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
func compareSlices(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestEFM_encodeQuattuordecuple(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    uint128.Uint128
		wantErr bool
	}{
		{
			name:  "valid input 8 bytes",
			input: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
			want:  uint128.Uint128{H: 145137984342281, L: 2486275223393945860},
		},
		{
			name:    "input less than 8 bytes",
			input:   []byte{0x01, 0x23, 0x45},
			wantErr: true,
		},
		{
			name:  "input with all 0xFF",
			input: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
			want:  uint128.Uint128{H: 35495776224392, L: 1306123611396147218},
		},
		{
			name:  "input with all 0x00",
			input: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			want:  uint128.Uint128{H: 79719458703378, L: 2326251190641758752},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()
			got, _ := e.encodeQuattuordecuple(tt.input)
			if !tt.wantErr && got != tt.want {
				t.Errorf("EFM.encodeQuattuordecuple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "returns non-nil pointer",
		},
		{
			name: "returns unique instances",
		},
		{
			name: "initializes valid internal state",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "returns non-nil pointer":
				obj := New()
				if obj == nil {
					t.Errorf("New() returned nil, want non-nil")
				}
			case "returns unique instances":
				obj1 := New()
				obj2 := New()
				if obj1 == obj2 {
					t.Errorf("New() returned the same instance for multiple calls, want unique instances")
				}
			case "initializes valid internal state":
				obj := New()
				if obj == nil {
					t.Fatalf("New() returned nil, can't check internal state")
				}
				if len(obj.encodeMapping) == 0 || len(obj.decodeMapping) == 0 {
					t.Errorf("New() did not initialize valid internal state")
				}
			}
		})
	}
}

func TestEFM_Encode(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "valid data encoding",
			input:   []byte("12345678"),
			wantErr: false,
		},
		{
			name:    "data length not divisible by 8",
			input:   []byte("12345"),
			wantErr: true,
		},
		{
			name:    "empty data",
			input:   []byte{},
			wantErr: false,
		},
		{
			name:    "large data divisible by 8",
			input:   []byte("1234567890ABCDEF1234567890ABCDEF"),
			wantErr: false,
		},
		{
			name:    "exactly 8 bytes",
			input:   []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			wantErr: false,
		},
		{
			name:    "data with null bytes",
			input:   []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			wantErr: false,
		},
		{
			name:    "special characters",
			input:   []byte("!@#$%^&*"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()
			_, err := e.Encode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("EFM.Encode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
