package util

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
)

type Enc int

const (
	Bin = iota
	Hex
	Dec
)

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

func PadWithString(msg, pad string, width int) string {
	minWidth := len(msg) + 2
	if minWidth > width {
		return ""
	}
	padCountLeft := (width - minWidth) / 2
	padCountRight := width - minWidth - padCountLeft
	var output string
	output += strings.Repeat(pad, padCountLeft)
	output += " " + msg + " "
	output += strings.Repeat(pad, padCountRight)
	return output
}
