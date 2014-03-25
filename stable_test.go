package stable

import (
	"encoding/csv"
	"fmt"
	"strings"
	"testing"
)

func TestErrorsWithoutRows(t *testing.T) {
	table := Table([][]string{})
	r := table.Select([]string{})

	if fmt.Sprintf("%v", r) != "[]" {
		t.Errorf("%v is not an empty Table", r)
	}

	table = Table([][]string{{}})
	r = table.Select([]string{})

	if fmt.Sprintf("%v", r) != "[[]]" {
		t.Errorf("%v is not an empty Table", r)
	}
}

func TestHeaderPositions(t *testing.T) {

	sample := `header a,header b,header c
a0,b0,c0
a1,b1,c1`

	doc, _ := csv.NewReader(strings.NewReader(sample)).ReadAll()

	f := Table(doc)
	f.setHeaders([]string{"header a", "header c"})
}

func ExampleTable_Select() {
	txt := "Column A,Column B\n" +
		"A1,B1\n" +
		"A2,B2"

	csvReader := csv.NewReader(strings.NewReader(txt))
	table, err := csvReader.ReadAll()

	if err != nil {
		panic(err)
	}

	f := Table(table)

	filtered := f.Select([]string{"Column B"})
	fmt.Printf("%v", filtered)

	// Output:
	// [[Column B] [B1] [B2]]
}

func ExampleTable_Merge() {
	tA := [][]string{
		{"Col A", "Col B"},
		{"A1", "B1"},
		{"", "B2"},
	}

	tB := [][]string{
		{"Col C"},
		{"C1"},
		{"C2"},
		{"C3"},
	}

	left := Table(tA)
	right := Table(tB)

	merged := left.Merge(right)
	fmt.Printf("%v", merged)

	// Output:
	// [[Col A Col B Col C] [A1 B1 C1] [ B2 C2] [  C3]]

}

// func ExampleTable_Diff() {
// }

// func TestDiffLeft(t *testing.T) {
// }

// func TestDiffRight(t *testing.T) {
// }

// func ExampleTable_Join() {
// }

// func ExampleTable_Rename() {
// }
