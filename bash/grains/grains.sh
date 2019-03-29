#!/usr/bin/env bash


MAX=64
MIN=1
BASE=2

# string comparison means you have to quote the args
if [[ "$1" == "total" ]]; then
    echo 2^64 | bc
    exit 0
fi

# number comparison with things like -eq -gt -lte
if [[ $MIN -gt $1 ]] || [[ $1 -gt $MAX ]] ; then
   echo "Error: invalid input" >&2;
   exit 1
fi

EXPONENT=$(($1-1))
echo "$(($BASE**$EXPONENT))"
exit 0
