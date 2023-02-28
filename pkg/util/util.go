package util

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/oklog/ulid"
	"strings"
)

type Enc int

const (
	Bin = iota
	Hex
	Dec
)

func ParseUlidString(id string) (ulid.ULID, bool) {
	val, err := ulid.ParseStrict(strings.TrimSpace(id))
	return val, err == nil
}

func BytesToString(b []byte, enc Enc) string {
	var output string
	n := len(b)
	switch enc {
	case Bin:
		for i, v := range b {
			output += fmt.Sprintf("%08b", uint8(v))
			if i < n-1 {
				output += " "
			}
		}
	case Hex:
		h := hex.EncodeToString(b)
		for i := 0; i < len(h)-1; i += 2 {
			output += h[i:i+2] + " "
		}
	case Dec:
		switch n {
		case 2:
			output = fmt.Sprintf("%v", binary.BigEndian.Uint16(b[:]))
		case 4:
			output = fmt.Sprintf("%v", binary.BigEndian.Uint32(b[:]))
		}
	}
	return output
}

func PadWithChar(msg, left, right, pad string, width int) string {
	minWidth := len(msg) + len(left) + len(right) + 2
	if minWidth > width {
		return ""
	}
	padCount := (width - minWidth) / 2
	var output string
	if left != "" {
		output += left
	}
	output += strings.Repeat(pad, padCount)
	output += " " + msg + " "
	if len(msg)%2 != 0 {
		padCount++
	}
	output += strings.Repeat(pad, padCount)
	if right != "" {
		output += right
	}
	return output
}
