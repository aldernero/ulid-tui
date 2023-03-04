package util

import (
	"github.com/oklog/ulid"
	"strings"
	"time"
)

type Field int

const (
	T1 Field = iota
	T2
	E1
	E2
	E3
)

type ULID struct {
	id ulid.ULID
	t1 [4]byte // 32 bit uint time high
	t2 [2]byte // 16 bit uint time low
	e1 [2]byte // 16 bit uint entropy
	e2 [4]byte // 32 bit uint entropy
	e3 [4]byte // 32 bit uint entropy
}

func NewULID(id ulid.ULID) ULID {
	var newULID ULID
	newULID.id = id
	bytes, err := id.MarshalBinary()
	if err != nil {
		panic(err)
	}
	copy(newULID.t1[:], bytes[:4])
	copy(newULID.t2[:], bytes[4:6])
	copy(newULID.e1[:], bytes[6:8])
	copy(newULID.e2[:], bytes[8:12])
	copy(newULID.e3[:], bytes[12:16])
	return newULID
}

func (u ULID) ToString(f Field, e Enc) string {
	switch f {
	case T1:
		return BytesToString(u.t1[:], e)
	case T2:
		return BytesToString(u.t2[:], e)
	case E1:
		return BytesToString(u.e1[:], e)
	case E2:
		return BytesToString(u.e2[:], e)
	case E3:
		return BytesToString(u.e3[:], e)
	}
	return ""
}

func UlidTimes(id ulid.ULID) (string, string, uint64) {
	ms := id.Time()
	local := time.UnixMilli(int64(ms)).Local().Format(time.RFC3339)
	utc := time.UnixMilli(int64(ms)).UTC().Format(time.RFC3339)
	return utc, local, ms
}

func ParseUlidString(id string) (ulid.ULID, bool, string) {
	val, err := ulid.ParseStrict(strings.TrimSpace(id))
	var reason string
	if err != nil {
		switch err {
		case ulid.ErrDataSize:
			reason = "Too short" // input is limited to 26 chars
		case ulid.ErrInvalidCharacters:
			reason = "Invalid characters"
		case ulid.ErrBigTime:
			reason = "Too far in the future"
		default:
			reason = "Invalid ULID"
		}
	}
	return val, err == nil, reason
}
