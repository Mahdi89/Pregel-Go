#!/bin/bash

# set Pregel_Project_root env variable
# to DIR where the project is cloned to
 
root=$Pregel_Project_root
plot_path=$root/plot
conf_file=config.go
out_file=bench.out

cd $plot_path && go build plot.go && ./plot $1
eog "$out_file.png"

