NAME := Pregel API
DESC := Pregel-based graph processing suite

.PHONY: install clean bench plot

install:
- go get -t github.com/skelterjohn/go.matrix

test: test_pagerank

test_pagerank: 
	cd ./pagerank/ && go test

bench: 
	cd ./bin/ && ./bench.sh 3 4

clean: 	
	cd ./bin/ && ./cleanup.sh
plot: clean bench
	cd ./bin/ && ./plot.sh "../pagerank/bench.out" 3
