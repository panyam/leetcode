#!/bin/sh

 while true; do
  clear
  go test -v ./...
  fswatch  -o ../ | echo "Files changed, re-testing..."
  sleep 1
 done
