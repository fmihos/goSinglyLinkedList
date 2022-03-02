package singlylinkedlist

import (
	"testing"
)

var list List

func TestInsertFront(t *testing.T) {
	t.Run("Insert to Head of list", func(t *testing.T) {
		list.InsertFront(5)
		list.InsertFront(6)

		expectedSize := 2
		expectedFirst := 6
		expectedSecond := 5

		if size := list.Size(); size != expectedSize {
			t.Errorf("Expected %d elements and got %d", expectedSize, size)
		}

		if firstNode := list.GetAt(0); (*firstNode).key != expectedFirst {
			t.Errorf("Expected %d as first element and got %d", expectedFirst, (*firstNode).key)
		}

		if secondNode := list.GetAt(1); (*secondNode).key != expectedSecond {
			t.Errorf("Expected %d as second element and got %d", expectedSecond, (*secondNode).key)
		}
	})
}
