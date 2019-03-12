#!/bin/bash

if [ -n "$(gofmt -l .)" ]; then
  echo "Run gofmt on your code:"
  gofmt -d .
  exit 1
fi
