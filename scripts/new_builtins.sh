#!/usr/bin/env bash

for i in *
do
if [[ $(wc -l $i/$i.go | awk '{print $1}') -gt 3 ]]; then
echo $i
fi
done
