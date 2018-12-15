#!/bin/bash

# set Pregel_Project_root env variable
# to DIR where the project is cloned to
 
root=$Project_root
conf_file="config.go"
out_file="bench.out"

# TODO allow iteration over 
# SUPER_STEPS and CONN_DEGREE too
cd $root/pagerank && rm -rf $out_file $conf_file
for value in $(seq 1 $1)
do

	cat > $conf_file <<- EOM
	package pagerank

	// Set high-level meta parameters
	const SUPER_STEPS = 100
	const CONN_DEGREE = $2
	const NUM_VERTEX = $(($value * 5))
	EOM

	go test -bench=. -benchmem >> $out_file
done
