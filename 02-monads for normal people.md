#Monads for Normal People
##Dustin Getz

Unit & bind are base of a monad, declare a standard for datatypes so that business
can be chained and separated for the actual program

Unit defines the datatype format (tuple of num, something for example)

Bind then establishes how to handle the "nuances" of the unit (test if None if the
format is a Maybe monad, etc), which allows for chains to not break if an error was
encountered (so bind handles checking for failure and continuing vs all shit blowing
up)

With unit & bind, you can then cleanly apply chain/pipe, map, etc.

Monads then define the namespace in how everything should look and exist (advances
the default datastructure without breaking the language) so that all ints are going
to look like (x, is_error, error_msg) or whatever which you can then define functions
to convey the states (so def success(val): return (val, False, "") )
