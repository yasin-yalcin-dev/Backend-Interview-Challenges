/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package stackqueue

import "fmt"

// RunExample demonstrates the stack and queue implementations
func RunExample() {
	fmt.Println("Stack & Queue Example:")
	fmt.Println("---------------------")

	// Stack Examples
	fmt.Println("\nStack Examples:")
	fmt.Println("1. Slice Implementation:")
	demonstrateStack(NewSliceStack())

	fmt.Println("\n2. Linked List Implementation:")
	demonstrateStack(NewLinkedStack())

	// Queue Examples
	fmt.Println("\nQueue Examples:")
	fmt.Println("1. Slice Implementation:")
	demonstrateQueue(NewSliceQueue())

	fmt.Println("\n2. Linked List Implementation:")
	demonstrateQueue(NewLinkedQueue())

	// Practical Applications
	fmt.Println("\nPractical Applications:")

	// Stack application: Parentheses matching
	fmt.Println("\n1. Using Stack for Parentheses Matching:")
	expressions := []string{
		"()",
		"(())",
		"(()){}[]",
		"({[]})",
		"({)}",
		"(((",
	}

	for _, expr := range expressions {
		isBalanced := checkBalancedParentheses(expr)
		fmt.Printf("Expression '%s' is balanced: %v\n", expr, isBalanced)
	}

	// Queue application: Print Queue
	fmt.Println("\n2. Using Queue for Print Queue Simulation:")
	simulatePrintQueue()
}

// demonstrateStack demonstrates basic stack operations
func demonstrateStack(stack Stack) {
	fmt.Println("- Pushing items: 10, 20, 30")
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Printf("- Stack size: %d\n", stack.Size())

	item, err := stack.Peek()
	if err == nil {
		fmt.Printf("- Top item (peek): %v\n", item)
	}

	fmt.Println("- Popping all items:")
	for !stack.IsEmpty() {
		item, err := stack.Pop()
		if err == nil {
			fmt.Printf("  Popped: %v\n", item)
		}
	}

	fmt.Printf("- Is stack empty? %v\n", stack.IsEmpty())
}

// demonstrateQueue demonstrates basic queue operations
func demonstrateQueue(queue Queue) {
	fmt.Println("- Enqueuing items: 10, 20, 30")
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Printf("- Queue size: %d\n", queue.Size())

	item, err := queue.Front()
	if err == nil {
		fmt.Printf("- Front item: %v\n", item)
	}

	fmt.Println("- Dequeuing all items:")
	for !queue.IsEmpty() {
		item, err := queue.Dequeue()
		if err == nil {
			fmt.Printf("  Dequeued: %v\n", item)
		}
	}

	fmt.Printf("- Is queue empty? %v\n", queue.IsEmpty())
}

// checkBalancedParentheses checks if an expression has balanced parentheses
func checkBalancedParentheses(expr string) bool {
	stack := NewStack()

	for _, char := range expr {
		switch char {
		case '(', '{', '[':
			stack.Push(char)
		case ')':
			item, err := stack.Pop()
			if err != nil || item.(rune) != '(' {
				return false
			}
		case '}':
			item, err := stack.Pop()
			if err != nil || item.(rune) != '{' {
				return false
			}
		case ']':
			item, err := stack.Pop()
			if err != nil || item.(rune) != '[' {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

// PrintJob represents a document to be printed
type PrintJob struct {
	ID       int
	Document string
	Pages    int
}

// simulatePrintQueue simulates a printer queue
func simulatePrintQueue() {
	printQueue := NewQueue()

	// Add print jobs
	printQueue.Enqueue(PrintJob{1, "Report.pdf", 5})
	printQueue.Enqueue(PrintJob{2, "Invoice.pdf", 2})
	printQueue.Enqueue(PrintJob{3, "Letter.doc", 1})

	fmt.Printf("Print queue has %d jobs\n", printQueue.Size())

	// Process print jobs
	for !printQueue.IsEmpty() {
		item, _ := printQueue.Dequeue()
		job := item.(PrintJob)
		fmt.Printf("Printing Job #%d: %s (%d pages)\n", job.ID, job.Document, job.Pages)
	}

	fmt.Println("All print jobs completed")
}
