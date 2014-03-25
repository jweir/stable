/*
Package stable allows merging, diffing and joining tables of string with type `[][]string`.

These methods and functions are useful for CSV, or similiar data. This package assumes the first row of data is the header.

This is a work in progress and not production ready.
*/
package stable

type Table [][]string

// Select returns a table with only columns having a headers that match the
// name of the columns given.
//
// A new, and empty, column will be created for any header which does not exist.
func (f Table) Select(columns []string) Table {
	indices := f.setHeaders(columns)

	out := [][]string{}

	for _, row := range f {
		nrow := []string{}
		for _, i := range indices {
			nrow = append(nrow, row[i])
		}
		out = append(out, nrow)
	}

	return Table(out)
}

// Merge combines two Tables into one
func (left Table) Merge(right Table) Table {
	out := Table{}

	for i := 0; i < rowCount(left, right); i++ {
		l := getRow(left, i)
		r := getRow(right, i)
		newRow := append(l, r...)
		out = append(out, newRow)
	}

	return out
}

// returns the row, or if there is no row
// a row padded with empty cells
func getRow(t Table, index int) []string {
	if len(t) > index {
		return t[index]
	} else {
		var empty []string
		for n := 0; n < len(t[0]); n++ {
			empty = append(empty, "")
		}
		return empty
	}
}

func rowCount(a, b Table) int {
	al := len(a)
	bl := len(b)

	if al >= bl {
		return al
	} else {
		return bl
	}
}

// Headers returns the slice of header names
func (f Table) headers() []string {
	if len(f) > 0 {
		return f[0]
	} else {
		return []string{}
	}
}

func (f Table) write(row []string) error {
	f = append(f, row)
	return nil
}

// return a set of indexes to keep
func (f *Table) setHeaders(columns []string) []int {
	var indices []int

	for i, key := range f.headers() {
		if has(key, columns) == true {
			indices = append(indices, i)
		}
	}

	return indices
}

func has(s string, set []string) bool {
	for _, v := range set {
		if v == s {
			return true
		}
	}

	return false
}
