package lister

import (
	"strconv"
	"testing"
)

func TestList(t *testing.T) {
	list := NewList()

	if list.Size() != 0 {
		t.Error("NewList l with empty paramaters is not l.Size() = 0")
	}
	list.Show()
	for i := range 10 {
		ch := strconv.Itoa(65 + (i % 26))
		list.Append(ch)
		list.Show()
	}
}
