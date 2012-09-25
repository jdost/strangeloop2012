#Using Haskell Type Signatures as a Functional Design Notation
##Michael Feathers

Haskell type signatures give an insight into what the functional will do

Informs how you can tie two functions together

Use the signature to help with chaining actions, you get a list of strings, you will
get out a single string, then you feed the single string in and get a tuple, etc.

Once you know the break down of the process of types, you can then fill in the
transformations with functions (and determine the names) and understand the planning
of each function

Not to be used as a strict standard, just a brain storming idea, not meant to be
verbose but instead aid in personal development

Use a full convention in naming, (uses `words` from Haskell to break lines into words
so uses naming for functions on what the function will produce (and not the action
or 'verb' of the function like 'breakIntoWords'))
