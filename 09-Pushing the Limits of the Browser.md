#Pushing the Limits of Web Browsers
##Lars Bak

###Why Language Based VMs
* Platform independent execution
* Sandboxing
* Optimiziations can take place at runtime
* Debugging is possible in production
* Loading of third party code at runtime
* When VM gets faster: programs run faster naturally, frees programmers to innovate

"Javascript is a weird language [...] At least it will be fun to make it fast"

###Dart
* Class-based single inheritance
* Interfaces with default implementation
* Optional static types
* Real lexical scoping
* Single-threaded

###Dart designed for VM
* Straight forward semantics
* Simple object model
* No class initialization
* Applications are declared

###Dart optional type system
* Used for specifying programmer intent
* Types can be introduced gradually
* Adding types to fields, variables, and method signatures does not change behavior
* Allows implicit down casting
* Verified only in checked mode
* Type checker only issues warnings

Dart source works in VM or translates into JS
Also can create a snapshot from source and run that in a VM (10x faster than using source)

###Dart SDK
* Language spec
* Libraries
* VM
* Translator to JS
* IDE
* Integration with Chromium

###Things Learned
* Always start with a small team
* Focus on solving the hardest problem first
* Competitive situation fuels motivation
* Only hire people that are smarter
* Open source projects are great: help industry, keep your work honest
* Track performance from day one (benchmarks on all revisions on all platforms)
* Testing must be run from day one (tests on all revisions on all platforms)

###Summary
* Speed will continue to drive innovation
* Programmer productivity will be key as web applications get larger
