#!/bin/bash

for dir in */; do
  if [ -d "$dir" ]; then
    if [ -f "${dir}demo.tape" ]; then
      # Check if demo.tape contains "sudo"
      if grep -q "sudo" "${dir}demo.tape"; then
        echo "Skipping demo in ${dir} because it contains 'sudo'."
      else
        echo "Running demo in ${dir}"
        (cd "$dir" && vhs demo.tape)
      fi
    else
      echo "No demo.tape found in ${dir}"
    fi
  fi
done

echo "Finished generating demos."