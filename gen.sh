#!/bin/bash
DELTA=.02
for i in `seq $DELTA $DELTA 1` ; do
  go run src/*.go -p=$i -dir=imgs &
done

for job in `jobs -p`
do
  wait $job
done

convert imgs/*.png imgs/`date +%s`.gif
rm imgs/*.png
