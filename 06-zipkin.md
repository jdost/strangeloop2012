#Zipkin: Distributed Tracing Framework
##Johan Oskarsson

Twitter's tracing system
OSS
Uses sampling to decrease overhead
Leverages the Finagle library (Twitter's) for RPC in the JVM (Scala and Java)
Minimizes the size used for tracking but just using 3 i64s to identify the trace
Ties in quickly with various services
