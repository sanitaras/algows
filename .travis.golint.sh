#!/bin/bash

if [ -n "$(golint .)" ]; then
  echo "Run golint on your code"
  golint .
  exit 1
fi
