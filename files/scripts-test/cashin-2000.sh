#!/bin/bash

for i in {1..2000}
do
  echo "Cashin number ${i}"
  ./cashin.sh &
done

echo "Finish"

