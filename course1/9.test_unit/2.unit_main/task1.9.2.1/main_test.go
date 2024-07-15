package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	oldOut := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("failed create pipe: %s", err)
	}

	os.Stdout = w

	main()
	w.Close()

	os.Stdout = oldOut

	var out bytes.Buffer
	out.ReadFrom(r)
	expected := "Hello world!\n"
	if out.String() != expected {
		t.Errorf("want: %s\ngot: %s", expected, out.String())
	}
}
