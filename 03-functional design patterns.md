#Functional Design Patterns
##Stuart Sierra

###State/Event Pattern
* State is defired from previous state + input
* Many diverse inputs
* Need to recover past astates
* Need to visualize intermediate states
* Takes the current state, the input, and updates the state
* Allows to recreate a state by storing the inputs (replay/recreate)
* Con: tons of data structures sitting in memory

###Consequences Pattern
* Every input can trigger multiple events
* Generated events cause state changes
* Need to visualize the intermediate states
* Do not compose naturally
* Have to update state in between

###Accumulator Pattern
* Large collection of inputs (maybe larger than memory)
* Small or scalar result
* Lazy sequences (map, filter, etc)
* reduce is the universal accumulator (turns a ton of data into something small)

###MapReduce
* Input is linear, maybe larger than 1 disk
* Disks are slow and local to one machine
* Networks are slow
* Not quite map and reduce
* Will become less important as hardware improves (SSD, faster networks)

###Reduce/Combine Pattern
* Input is tree-like
* Divide & Conquer approach
* Combining intermediate results is *associative*

###Recursive Expansion Pattern
* Build up result out of primitives
* Build abstractions in layers
* Recurse until no more work left to do
* Like macros in Lisp

###Pipeline Pattern
* Process with many discrete steps
* Similar "shape" of data at each step (map or record)
* Only one execution path
* Hard to escape (rolling ball) without throwing an exception

###Wrapper Pattern
* Process with many discrete steps
* One main execution path
* Possible branch at each step
* Stacks of functions that do things before/after to an input and then call the next function in the combination set

###Token Pattern
* May need to cancel an operation
* Operation itself is not an identity
* Gives access to a task that is opaque and controls its operation (cancel a scheduled task)

###Observer Pattern
* Register an observer function on a stateful container

###Strategy Pattern
* Many processes with simliar structure
* Need extension points for future variations
