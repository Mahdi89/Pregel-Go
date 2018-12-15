NAME := Pregel API
DESC := Pregel-based graph processing suite

.PHONY: test clean bench plot

test: test_pagerank

test_pagerank: 
	cd ./pagerank/ && go test

bench: 
	cd ./bin/ && ./bench.sh 3 4

clean: 	
	cd ./bin/ && ./cleanup.sh
plot: clean bench
	cd ./bin/ && ./plot.sh "../pagerank/bench.out" 3
