#!/bin/bash

if [ -n "$(go vet .)" ]; then
  echo "Run go vet on your code"
  go vet .
  exit 1
fi
