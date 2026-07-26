package util

import (
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
)

func TestNewULIDAndToString(t *testing.T) {
	id := ulid.MustParseStrict("01ARZ3NDEKTSV4RRFFQ69G5FAV")
	u := NewULID(id)

	tests := []struct {
		field Field
		want  string
	}{
		{T1, "01 56 3e 3a"},
		{T2, "b5 d3"},
		{E1, "d6 76"},
		{E2, "4c 61 ef b9"},
		{E3, "93 02 bd 5b"},
	}
	for _, tt := range tests {
		got := u.ToString(tt.field, Hex)
		if got != tt.want {
			t.Errorf("ToString(field=%v, Hex) = %q, want %q", tt.field, got, tt.want)
		}
	}
}

func TestToStringUnknownField(t *testing.T) {
	u := NewULID(ulid.MustParseStrict("01ARZ3NDEKTSV4RRFFQ69G5FAV"))
	if got := u.ToString(Field(99), Hex); got != "" {
		t.Errorf("ToString(unknown field) = %q, want empty string", got)
	}
}

func TestUlidTimes(t *testing.T) {
	id := ulid.MustParseStrict("01ARZ3NDEKTSV4RRFFQ69G5FAV")
	utc, _, ms := UlidTimes(id)
	wantMs := uint64(1469922850259)
	if ms != wantMs {
		t.Errorf("UlidTimes ms = %d, want %d", ms, wantMs)
	}
	wantUTC := time.UnixMilli(int64(wantMs)).UTC().Format(time.RFC3339)
	if utc != wantUTC {
		t.Errorf("UlidTimes utc = %q, want %q", utc, wantUTC)
	}
}

func TestUlidTimesZero(t *testing.T) {
	var id ulid.ULID
	_, _, ms := UlidTimes(id)
	if ms != 0 {
		t.Errorf("UlidTimes(zero ULID) ms = %d, want 0", ms)
	}
}

func TestParseUlidString(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantValid  bool
		wantReason string
	}{
		{"valid", "01ARZ3NDEKTSV4RRFFQ69G5FAV", true, ""},
		{"valid with whitespace", "  01ARZ3NDEKTSV4RRFFQ69G5FAV  ", true, ""},
		{"too short", "01ARZ3NDEK", false, "Too short"},
		{"invalid characters", "!!ARZ3NDEKTSV4RRFFQ69G5FAV", false, "Invalid characters"},
		// ulid.ParseStrict reports an overflowing timestamp as ErrOverflow,
		// not ErrBigTime (the latter is only returned by construction
		// functions like ulid.New), so this falls through to the default reason.
		{"overflowing timestamp", "ZZZZZZZZZZZZZZZZZZZZZZZZZZ", false, "Invalid ULID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, valid, reason := ParseUlidString(tt.input)
			if valid != tt.wantValid {
				t.Errorf("ParseUlidString(%q) valid = %v, want %v", tt.input, valid, tt.wantValid)
			}
			if reason != tt.wantReason {
				t.Errorf("ParseUlidString(%q) reason = %q, want %q", tt.input, reason, tt.wantReason)
			}
		})
	}
}
