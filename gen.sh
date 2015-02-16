#!/bin/bash
go install github.com/theepicsnail/gofractal/flame || exit 1
HEIGHT=400
WIDTH=400
ITERATIONS=1000000
FILE="imgs/img%.4f.png"
DELTA=.02
for P in `seq $DELTA $DELTA 1` ; do
  flame -height=$HEIGHT -width=$WIDTH -iterations=$ITERATIONS -p=$P -file=$FILE
done

for job in `jobs -p`
do
  wait $job
done

convert imgs/*.png imgs/latest.gif
cp imgs/latest.gif imgs/`date +%s`.gif
rm imgs/*.png
