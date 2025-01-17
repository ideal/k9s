package render

import (
	"reflect"
)

// DeltaRow represents a collection of row detlas between old and new row.
type DeltaRow []string

// NewDeltaRow computes the delta between 2 rows.
func NewDeltaRow(o, n Row, excludeLast bool) DeltaRow {
	deltas := make(DeltaRow, len(o.Fields))
	// Exclude age col
	oldFields := o.Fields[:len(o.Fields)-1]
	if !excludeLast {
		oldFields = o.Fields[:len(o.Fields)]
	}
	for i, old := range oldFields {
		if old != "" && old != n.Fields[i] {
			deltas[i] = old
		}
	}

	return deltas
}

// Diff returns true if deltas differ or false otherwise.
func (d DeltaRow) Diff(r DeltaRow, ageCol int) bool {
	if len(d) != len(r) {
		return true
	}

	if ageCol < 0 || ageCol >= len(d) {
		return !reflect.DeepEqual(d, r)
	}

	if !reflect.DeepEqual(d[:ageCol], r[:ageCol]) {
		return true
	}
	if ageCol+1 >= len(d) {
		return false
	}

	return !reflect.DeepEqual(d[ageCol+1:], r[ageCol+1:])
}

// Customize returns a subset of deltas.
func (d DeltaRow) Customize(cols []int, out DeltaRow) {
	if d.IsBlank() {
		return
	}
	for i, c := range cols {
		if c < 0 {
			continue
		}
		if c < len(d) && i < len(out) {
			out[i] = d[c]
		}
	}
}

// IsBlank asserts a row has no values in it.
func (d DeltaRow) IsBlank() bool {
	if len(d) == 0 {
		return true
	}

	for _, v := range d {
		if v != "" {
			return false
		}
	}

	return true
}

// Clone returns a delta copy.
func (d DeltaRow) Clone() DeltaRow {
	res := make(DeltaRow, len(d))
	for i, f := range d {
		res[i] = f
	}

	return res
}
