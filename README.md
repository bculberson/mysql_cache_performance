# mysql_cache_performance
Project to performance benchmark using mysql reader as a cache vs redis

# to run:

* docker-machine create mysqlbench
* eval "$(docker-machine env mysqlbench)"
* docker-compose up mysqlbench

Example output:

> bench_1 | testing: warning: no tests to run

> bench_1 | PASS

> bench_1 | BenchmarkMySqlGet	   20000	     79181 ns/op

> bench_1 | BenchmarkRedisGet	   50000	     30185 ns/op

> bench_1 | ok  	_/code	4.311s
