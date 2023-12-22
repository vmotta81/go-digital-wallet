#!/bin/bash

for i in {1..10000}
do
  echo "Cashin number ${i}"
  ./cashin.sh "${i}" &
done

echo "Finish"

