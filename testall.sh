#!/bin/bash

for D in $(ls -1d sd*); do
  go test "github.com/gaorx/stardust4/$D"
done