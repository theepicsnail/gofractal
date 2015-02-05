#!/bin/bash
rm imgs/*.png
for i in `seq 0 .02 1` ; do
  go run src/*.go -p=$i -dir=imgs &
done

for job in `jobs -p`
do
  wait $job
done

convert imgs/*.png imgs/`date +%s`.gif
