#!/bin/bash

for dir in examples/*/; do
  if [ -d "$dir" ]; then
    if [ -f "${dir}demo.tape" ]; then
      echo "Running demo in ${dir}"
      (cd "$dir" && vhs demo.tape)
    else
      echo "No demo.tape found in ${dir}"
    fi
  fi
done

echo "Finished generating demos."
