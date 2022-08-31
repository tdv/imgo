#!/bin/bash

prefix_dir=components
components=(postgres redis service)

for i in ${components[*]}
do
  echo "Processing component \"$i\""
  echo "Entering directory \"$PWD/$prefix_dir/$i\""
  cd ./$prefix_dir/$i
  ./build.sh
  echo "Leaving directory \"$PWD\""
  cd ../../
  echo "The component \"$i\" is processed."
done
