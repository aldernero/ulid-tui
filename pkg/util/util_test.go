package util

import "testing"

func TestBytesToStringBin(t *testing.T) {
	got := BytesToString([]byte{0x01, 0xff}, Bin)
	want := "00000001 11111111"
	if got != want {
		t.Errorf("BytesToString(Bin) = %q, want %q", got, want)
	}
}

func TestBytesToStringHex(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want string
	}{
		{"two bytes", []byte{0xb5, 0xd3}, "b5 d3"},
		{"four bytes", []byte{0x01, 0x56, 0x3e, 0x3a}, "01 56 3e 3a"},
		{"single byte", []byte{0xab}, "ab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToString(tt.in, Hex)
			if got != tt.want {
				t.Errorf("BytesToString(Hex) = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBytesToStringDec(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want string
	}{
		{"two bytes", []byte{0x00, 0x01}, "1"},
		{"four bytes", []byte{0x00, 0x00, 0x00, 0x2a}, "42"},
		{"unsupported length", []byte{0x01}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToString(tt.in, Dec)
			if got != tt.want {
				t.Errorf("BytesToString(Dec) = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPadWithString(t *testing.T) {
	tests := []struct {
		name  string
		msg   string
		pad   string
		width int
		want  string
	}{
		{"even padding", "Time", "═", 12, "═══ Time ═══"},
		{"exact fit", "AB", "-", 4, " AB "},
		{"too narrow returns empty", "Hello", "-", 3, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PadWithString(tt.msg, tt.pad, tt.width)
			if got != tt.want {
				t.Errorf("PadWithString(%q, %q, %d) = %q, want %q", tt.msg, tt.pad, tt.width, got, tt.want)
			}
			if got != "" && len([]rune(got)) != tt.width {
				t.Errorf("PadWithString(%q, %q, %d) produced width %d, want %d", tt.msg, tt.pad, tt.width, len([]rune(got)), tt.width)
			}
		})
	}
}
