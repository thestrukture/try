package try


// Task abstracts passing code blocks
// into the try/finally statement
type Task func()

// Error abstracts passing code blocks
// to handle errors generated
type Error func(e interface{})

// TryStatement is the type returned
// on execution of package function Try..
// This enables chaining functions that compose
// the try/catch/finally statement.
type TryStatement struct {
	// Specifies the code block to run
	run Task
	// Specifes the code block to run
	// when an exception occurs
	errorHandler Error
}	
