#!/usr/bin/env bash

PERSON="you"
if [[ "$1" ]]; then
    PERSON="$1"
fi

echo "One for $PERSON, one for me."
exit 0
