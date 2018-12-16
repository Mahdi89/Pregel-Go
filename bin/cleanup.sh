#!/bin/bash
 
root=$Project_root
plot_path=$root/plot
conf_file=config.go
out_file=bench.out

cd $root && rm -rf $conf_file $out_file
cd $plot_path && rm -rf "$out_file.png" plot


