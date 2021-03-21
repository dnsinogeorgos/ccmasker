#!/usr/bin/env bash

go build .

time zcat testdata/*.gz | ./ccmasker -cpuprofile=cpu.prof > /dev/null
echo 'svg > cpu.svg' | go tool pprof cpu.prof 2&> /dev/null

time zcat testdata/*.gz | ./ccmasker -memprofile=mem.prof > /dev/null
echo 'svg > mem_space.svg' | go tool pprof -alloc_space mem.prof 2&> /dev/null
echo 'svg > mem_objects.svg' | go tool pprof -alloc_objects mem.prof 2&> /dev/null

#time zcat testdata/*.gz | ./ccmasker.py > /dev/null

#go tool pprof -http localhost:6060 cpu.prof
#go tool pprof -http localhost:6061 mem.prof
