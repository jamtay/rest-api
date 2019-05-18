run `go build && ./rest-api` from rest-api directory


Set up db by
 1. `brew install mongodb`
 2. `mongod` <- starts daemon
 3. `mongo` <- starts mongo shell
 4. `use movies_db` and `use people_db`  <- creates db (run from terminal window running mongo shell)
 5. `quit()` <- quits shell
 6. `ctrl-c` < quits daemon_
 
 
 
 **TODO**
 
 - Learn about imports (https://tour.golang.org/basics/10 - learn about imports and rename things properly)
 - find out how to do unit tests
 
 
 **Up To here in Go tutorial**
 https://tour.golang.org/concurrency/11
 **Things learnt**
 
 - No need for brackets in `for` and `if`
 - `for` loop without pre and post conditions is Gos `while` loop
 - Can declare a pre action in an `if` and the variable is only in scope for `if/else` statement (see `if` statement in `resource` directory regarding `err := `)
 - Gos switch statement does not automatically execute next case (`break` statement is always in each block (unlike `java/js`))
 - `defer` statement is executed at the end of the function.  Arguments are resolved immediately but statement is not executed until after function
 - `defer` statement execute in LIFO order if multiple in one function 
 - `a[low : high]` slicing an array includes the first element (low) but excludes the last (high)
 - Changing the elements of a slice modifies the corresponding elements of its underlying array.
 - When you don't have the number in the [] it is a slice literal that references the underlying array.  `q := []int{2, 3, 5, 7, 11, 13}` e.g
 - When slicing, you may omit the high or low bounds to use their defaults instead.  The default is zero for the low bound and the length of the slice for the high bound.
 - A slice has both a length and a capacity.  The length of a slice is the number of elements it contains.  The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
 - A method is just a function with a receiver argument.  See `movies_dao`
 - Methods with pointer receivers can modify the value to which the receiver points
 - An interface type is defined as a set of method signatures.  A value of interface type can hold any value that implements those methods.  e.g `type Abser interface {Abs() float64}`.  To declare something as type `Abser` the type created has implement a method called `Abs()`
 - Empty interfaces - https://tour.golang.org/methods/14 used to handle values of unknown types
 - Switch statement based on type of variable - https://tour.golang.org/methods/16
 - https://tour.golang.org/methods/17 - to string method for types and structs
 - Handling errors - https://tour.golang.org/methods/19.  Handle with if err!=nil
 - A goroutine is a lightweight thread managed by the Go runtime.
 - Channels - By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables. https://tour.golang.org/concurrency/2
 - Can loop over channels - https://tour.golang.org/concurrency/4.  Useful