# In-memory databases
## Michael Stonebraker, VoltDB

###Traditional Transaction Processing
* 1985, 1000 transactions/second (airplane tickets)
* ACID

###New TP
* MMO games require a lot of transactions
* High volume
* Sensor tagging (tagged objects being tracked)
* Eventually these sensors end with TP
* Electronic commerce: wall street trading, micro transactions, fraud detection

###In all cases (traditional, "New")
* ACID problem
* Workload is updates & queries
* Firehose
* Break traditional solutions (scale & latency)


1 TB is a "really big" DB ($50K for 1 TB main memory -- 64 GB across 16 servers)
Shore DBMS - 4% of performance is actual work (24% each for latching, locking, recovery, buffer pool)

###Faster
* Focus on overhead (better b-trees affects only 4% of path length)
* Get rid of ALL major sources of overhead (main memory deployment removes buffer pools)

###Options
* OldSQL (legacy RDBMS)
 * Old code lines (from 1980's)
 * "bloatware"
 * Spend all of their time on overhead, this is written into the large codebase
 * Reading: *The innovators dilemma" by Clayton Christenson
 * Longterm drift
* NoSQL (forget SQL and ACID for performance)
 * Give up traditional ACID and SQL
 * Compiler translates SQL into low level operations
 * 30 years of RDBMS Experience: Hard to beat compiler, Stored Procs are good (1 trip from app to DBMS)
 * Data consistency without ACID sucks
 * Hard to guarantee you won't need ACID later
* NewSQL (performance from new architecture)

###ACID needs
* Funds transfer
* Integrity constraints
* Multi-record state (both transactions happen or neither happens)

###VoltDB
* OSS
* 70x faster than "OldSQL"
* 5-7x faster than Cassandra as K-V DB
* Single threaded, no latching, no locking
