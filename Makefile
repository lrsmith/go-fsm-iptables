EXAMPLE_DIR=examples
DOT= dot -Tpng 

all: clean build example_one

examples: example_one

example_one:
	./go-fsm-iptables -f $(EXAMPLE_DIR)/data/example_one.output > $(EXAMPLE_DIR)/gv/example_one.gv
	$(DOT)  $(EXAMPLE_DIR)/gv/example_one.gv -o  $(EXAMPLE_DIR)/graphs/graph_one.png

example_clean:
	rm -f $(EXAMPLE_DIR)/gv/*
	rm -f $(EXAMPLE_DIR)/graphs/*

clean:
	rm go-fsm-iptables

build:
	go build
