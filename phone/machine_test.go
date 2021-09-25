package parser_test

import (
	"testing"

	parser "github.com/alextanhongpin/go-ragel-test"
	"github.com/google/go-cmp/cmp"
)

func TestCorrectValues(t *testing.T) {
	p := &parser.Phone{
		IntCode:  "1",
		AreaCode: "555",
		Number:   "2334567",
	}
	testMachine(t, p, "+1 (555) 2334567")
	testMachine(t, p, "+1(555)2334567")
}

func testMachine(t *testing.T, exp *parser.Phone, in string) {
	machine := parser.New()
	p, err := machine.Parse([]byte(in))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(exp, p); diff != "" {
		t.Fatal(diff)
	}
}
