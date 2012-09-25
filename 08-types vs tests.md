#Types vs Tests: An Epic Battle?
##Paul Snively & Amanda Laucher

###Assumptions
* REPL (Know how to use a REPL in our development)
* Logic
* Some languages are better suited for correct code
* Read vs Write (spend more time reading code)
* We are lazy

Michael Feathers defines legacy code as code without automated tests...

###Stereotypes
* It's easier to refactor with tests than types
* No tests = no trust
* Tests take a long time to run & types to compile
* Property based testing can replace unit testing
* Modular design only occurs with TDD
* I don't get errors that can be prevented by types
* Ivory Tower vs Hippies
* 100% coverage
* Typed code is verbose
* Types take too long
* Testing is for QA

###Kata
* used 3 "stories" to test approaches

###Amanda's Approach
* F#
* Types
* REPL play for algos
* Tests for validation
* Types save having to write tests for certain categories

###Paul's Approach
* Scala
* Tests first
* Used typing to make sure that account numbers are valid
* Based on checksums, compiler would fail for invalid accounts
* Test code is still code, with its own correctness, maintenance, etc (burden)

###Discussion
* Small codebase = little value for type system
* Types scale better
* Types have little value when talking with non technical end users
* Understanding requirements are the hardest part
* Rarely have the luxury of sample input/output

Tests are created based on examples
Types are built in, no need for examples to use
