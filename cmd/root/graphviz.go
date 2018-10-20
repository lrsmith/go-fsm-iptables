package cmd

// https://www.graphviz.org/doc/info/attrs.html
// https://renenyffenegger.ch/notes/tools/Graphviz/examples/index

import "fmt"

//	label="Input Format: rule_order: (prot, in, out, source, destination : special rules )";
//	labelloc=bottom;
//	labeljust=left;

var GraphvizHeader = `
digraph finite_state_machine {
	rankdir=LR;
	size="8,5"

	node [shape = doublecircle]; ACCEPT
	node [shape = doublecircle]; DROP
	node [shape = doublecircle]; MARK
	node [shape = circle];

	{ rank=same ACCEPT DROP MARK }

`

var GraphvizFooter = `
}
`

func PrintNode(state, newstate, input string) {

	// fmt.Printf("    \"%s\" -> \"%s\" [ label = \"%v\" ];\n", state, newstate, input)
	fmt.Printf("    \"%s\" -> \"%s\";\n", state, newstate)
}

func OutputGraphvizDot() {

	if verbose {
		fmt.Printf("Starting with chain: %s\n", startState)
	}

	fmt.Println(GraphvizHeader)
	FollowChain(startState)
	fmt.Println(GraphvizFooter)
}
