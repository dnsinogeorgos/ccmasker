#!/usr/bin/env bash

go build .

echo "cpu profiling on real data"
time zcat testdata/*.gz | ./ccmasker -cpuprofile=cpu.prof > /dev/null
echo 'svg > cpu.svg' | go tool pprof cpu.prof 2&> /dev/null

echo "heap profiling on real data"
time zcat testdata/*.gz | ./ccmasker -memprofile=mem.prof > /dev/null
echo 'svg > mem_space.svg' | go tool pprof -alloc_space mem.prof 2&> /dev/null
echo 'svg > mem_objects.svg' | go tool pprof -alloc_objects mem.prof 2&> /dev/null

#echo "timing ccmasker written in go (1.16)"
#time zcat testdata/*.gz | ./ccmasker > /dev/null
#echo
#echo "timing ccmasker written in python (3.9)"
#time zcat testdata/*.gz | ./ccmasker.py > /dev/null

#go tool pprof -http localhost:6060 cpu.prof
#go tool pprof -http localhost:6061 mem.prof
