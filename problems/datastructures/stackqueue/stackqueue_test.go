/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

// problems/datastructures/stackqueue/stackqueue_test.go
package stackqueue

import "testing"

// Test Stack implementations
func TestStackImplementations(t *testing.T) {
	// Test slice-based stack
	t.Run("SliceStack", func(t *testing.T) {
		testStackOperations(t, NewSliceStack())
	})

	// Test linked list-based stack
	t.Run("LinkedStack", func(t *testing.T) {
		testStackOperations(t, NewLinkedStack())
	})
}

// Test Queue implementations
func TestQueueImplementations(t *testing.T) {
	// Test slice-based queue
	t.Run("SliceQueue", func(t *testing.T) {
		testQueueOperations(t, NewSliceQueue())
	})

	// Test linked list-based queue
	t.Run("LinkedQueue", func(t *testing.T) {
		testQueueOperations(t, NewLinkedQueue())
	})
}

// testStackOperations tests the basic operations of a stack
func testStackOperations(t *testing.T, stack Stack) {
	// Test initial state
	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}

	if stack.Size() != 0 {
		t.Errorf("Expected size 0, got %d", stack.Size())
	}

	// Test peek and pop on empty stack
	_, err := stack.Peek()
	if err == nil {
		t.Error("Peek on empty stack should return error")
	}

	_, err = stack.Pop()
	if err == nil {
		t.Error("Pop on empty stack should return error")
	}

	// Test push
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	if stack.IsEmpty() {
		t.Error("Stack should not be empty after push")
	}

	if stack.Size() != 3 {
		t.Errorf("Expected size 3, got %d", stack.Size())
	}

	// Test peek
	item, err := stack.Peek()
	if err != nil {
		t.Errorf("Peek should not return error: %v", err)
	}

	if item != 30 {
		t.Errorf("Expected top item 30, got %v", item)
	}

	// Test pop
	item, err = stack.Pop()
	if err != nil {
		t.Errorf("Pop should not return error: %v", err)
	}

	if item != 30 {
		t.Errorf("Expected popped item 30, got %v", item)
	}

	if stack.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", stack.Size())
	}

	// Pop remaining items
	stack.Pop()
	item, _ = stack.Pop()

	if item != 10 {
		t.Errorf("Expected popped item 10, got %v", item)
	}

	if !stack.IsEmpty() {
		t.Error("Stack should be empty after popping all items")
	}
}

// testQueueOperations tests the basic operations of a queue
func testQueueOperations(t *testing.T, queue Queue) {
	// Test initial state
	if !queue.IsEmpty() {
		t.Error("New queue should be empty")
	}

	if queue.Size() != 0 {
		t.Errorf("Expected size 0, got %d", queue.Size())
	}

	// Test front and dequeue on empty queue
	_, err := queue.Front()
	if err == nil {
		t.Error("Front on empty queue should return error")
	}

	_, err = queue.Dequeue()
	if err == nil {
		t.Error("Dequeue on empty queue should return error")
	}

	// Test enqueue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	if queue.IsEmpty() {
		t.Error("Queue should not be empty after enqueue")
	}

	if queue.Size() != 3 {
		t.Errorf("Expected size 3, got %d", queue.Size())
	}

	// Test front
	item, err := queue.Front()
	if err != nil {
		t.Errorf("Front should not return error: %v", err)
	}

	if item != 10 {
		t.Errorf("Expected front item 10, got %v", item)
	}

	// Test dequeue
	item, err = queue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue should not return error: %v", err)
	}

	if item != 10 {
		t.Errorf("Expected dequeued item 10, got %v", item)
	}

	if queue.Size() != 2 {
		t.Errorf("Expected size 2 after dequeue, got %d", queue.Size())
	}

	// Dequeue remaining items
	queue.Dequeue()
	item, _ = queue.Dequeue()

	if item != 30 {
		t.Errorf("Expected dequeued item 30, got %v", item)
	}

	if !queue.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all items")
	}
}

// Test parentheses matching function
func TestCheckBalancedParentheses(t *testing.T) {
	testCases := []struct {
		expression string
		isBalanced bool
	}{
		{"()", true},
		{"(())", true},
		{"(()){}[]", true},
		{"({[]})", true},
		{"", true},
		{"(", false},
		{")", false},
		{"({)}", false},
		{"(()", false},
		{"]", false},
	}

	for _, tc := range testCases {
		result := checkBalancedParentheses(tc.expression)
		if result != tc.isBalanced {
			t.Errorf("For expression '%s', expected balance: %v, got: %v",
				tc.expression, tc.isBalanced, result)
		}
	}
}
