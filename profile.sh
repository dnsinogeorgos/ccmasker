#!/usr/bin/env bash

messageArray=(
  "this log message asdfasdfasdfasdf contains a asdfasdfasdfasdf asdfas PAN but also that asdfasdfasdfasdf asdfas PAN"
  "this log message asdfasdfasdfasdf contains a asdfasdfasdfasdf asdfas PAN but also that asdfasdfasdfasdf asdfas PAN"
  "this log message 53111111  111111 contains a 3611-2222  333+4444 123 PAN but also that 3611-2222  333+4444 123 PAN"
  "this log message 5311111111111111 contains a 3611-2222_3333+4444 123 PAN but also that 3611-2222_3333+4444 123 PAN"
)

go build .

#time for ((i = 0; i < 10000000; i++))
#do
#  echo "${messageArray[3]}"
#done | ./ccmasker > /dev/null

time for ((i = 0; i < 2500000; i++))
do
  echo "${messageArray[0]}"
  echo "${messageArray[1]}"
  echo "${messageArray[2]}"
  echo "${messageArray[3]}"
done | ./ccmasker -cpuprofile=cpu.prof > /dev/null
echo 'svg > cpu.svg' | go tool pprof cpu.prof 2&> /dev/null

time for ((i = 0; i < 2500000; i++))
do
  echo "${messageArray[0]}"
  echo "${messageArray[1]}"
  echo "${messageArray[2]}"
  echo "${messageArray[3]}"
done | ./ccmasker -memprofile=mem.prof > /dev/null
echo 'svg > mem_space.svg' | go tool pprof -alloc_space mem.prof 2&> /dev/null
echo 'svg > mem_objects.svg' | go tool pprof -alloc_objects mem.prof 2&> /dev/null

#go tool pprof -http localhost:6060 cpu.prof
#go tool pprof -http localhost:6061 mem.prof
