#!/bin/bash

for i in {1..10000}
do
  echo "Cashout number ${i}"
  ./cashout.sh "${i}" &
done

echo "Finish"

