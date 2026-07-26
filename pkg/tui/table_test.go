package tui

import (
	"strings"
	"testing"

	"github.com/aldernero/ulid-tui/pkg/util"
	"github.com/oklog/ulid/v2"
)

func TestGenHeaderContainsMessage(t *testing.T) {
	out := genHeader("Encoding", tableWidth)
	if !strings.Contains(out, "Encoding") {
		t.Errorf("genHeader output missing message: %q", out)
	}
}

func TestGenFooterContainsMessage(t *testing.T) {
	out := genFooter("4 bytes", tableWidth)
	if !strings.Contains(out, "4 bytes") {
		t.Errorf("genFooter output missing message: %q", out)
	}
}

func TestEmptyEncodingTableShowsPlaceholders(t *testing.T) {
	out := emptyEncodingTable()
	for _, want := range []string{t1Placeholder, t2Placeholder, e1Placeholder, e2Placeholder, e3Placeholder} {
		if !strings.Contains(out, want) {
			t.Errorf("emptyEncodingTable() missing placeholder %q", want)
		}
	}
}

func TestCreateEncodingTableShowsEncodedValues(t *testing.T) {
	id := ulid.MustParseStrict("01ARZ3NDEKTSV4RRFFQ69G5FAV")
	out := createEncodingTable(id, util.Hex)
	for _, want := range []string{"01 56 3e 3a", "b5 d3", "d6 76", "4c 61 ef b9", "93 02 bd 5b"} {
		if !strings.Contains(out, want) {
			t.Errorf("createEncodingTable(Hex) missing %q in output", want)
		}
	}
}

func TestCreateTimeTableWithValidULID(t *testing.T) {
	id := ulid.MustParseStrict("01ARZ3NDEKTSV4RRFFQ69G5FAV")
	out := createTimeTable(id)
	if !strings.Contains(out, "2016-07-30") {
		t.Errorf("createTimeTable() missing expected date: %q", out)
	}
	if !strings.Contains(out, "1469922850259") {
		t.Errorf("createTimeTable() missing expected epoch ms: %q", out)
	}
}

func TestCreateTimeTableWithZeroULID(t *testing.T) {
	var id ulid.ULID
	out := createTimeTable(id)
	if strings.Contains(out, "milliseconds") {
		t.Errorf("createTimeTable() for zero ULID should omit epoch value, got: %q", out)
	}
}

func TestCreateUlidStringBreakdownShort(t *testing.T) {
	out := createUlidStringBreakdown("01ARZ3ND")
	if !strings.Contains(out, "░") {
		t.Errorf("expected placeholder blocks for short input, got: %q", out)
	}
}

func TestCreateUlidStringBreakdownFull(t *testing.T) {
	out := createUlidStringBreakdown("01arz3ndektsv4rrffq69g5fav")
	if !strings.Contains(out, "01ARZ3NDEK") {
		t.Errorf("expected uppercased time portion, got: %q", out)
	}
	if !strings.Contains(out, "TSV4RRFFQ69G5FAV") {
		t.Errorf("expected uppercased entropy portion, got: %q", out)
	}
}
