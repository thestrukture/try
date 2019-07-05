package try


// Run initializes the try statement.
// The function passed is stored, to be ran
// later.
func Run(t Task) TryStatement {
	return TryStatement{ 
		run : t,
	}
}

// Throw abstracts writing error
// check statements. Throw will invoke panic
// when the parameter passed does not equal
// nil.
func Throw(e interface{}){
	if e != nil {
		panic(e)
	}
}

// Catch stores the error handler code block
// to the Try Statement.
func (t TryStatement) Catch(e Error) TryStatement {
	t.errorHandler = e
	return t
}

// Finally executes the statement and recovers from any 
// panics.
func (t TryStatement) Finally(task Task) {

	defer func() {
        if r := recover(); r != nil {
            t.errorHandler(r)
        }
    }()

	t.run()

	task()

}