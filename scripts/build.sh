#!/bin/bash

go build gotop.go

if [[ $? -ne 0 ]]; then
  echo "Build failed"
  exit 1
fi