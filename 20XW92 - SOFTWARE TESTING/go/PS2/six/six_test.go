package six

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	dll := &DoublyLinkedList{}

	// Test InsertAtEnd and PrintList
	dll.InsertAtEnd(1)
	dll.InsertAtEnd(2)
	dll.InsertAtEnd(3)
	expected := "1 2 3 "
	if got := dll.PrintList(); got != expected {
		t.Errorf("PrintList() = %v, want %v", got, expected)
	}

	// Test InsertAtBeginning
	dll.InsertAtBeginning(0)
	expected = "0 1 2 3 "
	if got := dll.PrintList(); got != expected {
		t.Errorf("PrintList() = %v, want %v", got, expected)
	}

	// Test DeleteByValue
	dll.DeleteByValue(2)
	expected = "0 1 3 "
	if got := dll.PrintList(); got != expected {
		t.Errorf("PrintList() = %v, want %v", got, expected)
	}

	// Test Deletion of Head
	dll.DeleteByValue(0)
	expected = "1 3 "
	if got := dll.PrintList(); got != expected {
		t.Errorf("PrintList() = %v, want %v", got, expected)
	}

	// Test Deletion of Tail
	dll.DeleteByValue(3)
	expected = "1 "
	if got := dll.PrintList(); got != expected {
		t.Errorf("PrintList() = %v, want %v", got, expected)
	}

	// Test Deletion of Non-existing Element
	dll.DeleteByValue(10)
	expected = "1 "
	if got := dll.PrintList(); got != expected {
		t.Errorf("PrintList() = %v, want %v", got, expected)
	}

	// Test Deletion of Last Element
	dll.DeleteByValue(1)
	expected = ""
	if got := dll.PrintList(); got != expected {
		t.Errorf("PrintList() = %v, want %v", got, expected)
	}
}
