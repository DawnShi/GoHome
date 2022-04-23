#! /bin/bash

# go tool pprof -http=:8090 cpu-large.out
# go tool pprof -http=:8090 mem-large.out
go tool pprof -http=:8090 cpu-small.out