package cmd

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func parseIPTables(filename string) {

	commentRegex := regexp.MustCompile(`\/\*.*\*\/`)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Found start of a new chain
		if strings.HasPrefix(line, "Chain") {
			chainName := strings.Fields(line)[1]

			// Create a struct for the chain
			IPChains[chainName] = &IPChain{Name: chainName, Rules: []ChainRule{}}

			// Read the Chain's rules
			for scanner.Scan() {

				var extendedFields string
				ruleLine := commentRegex.ReplaceAllString(scanner.Text(), "")

				// Reached empty line. End of Chain.
				if len(ruleLine) == 0 {
					break
				}
				fields := strings.Fields(ruleLine)
				if fields[0] == "pkts" {
					continue
				}
				if len(fields) > 9 {
					extendedFields = strings.Join(fields[9:], " ")
				}

				// Add the rule to the Chain's struct
				IPChains[chainName].AddRule(ChainRule{Target: fields[2],
					Prot: fields[3], In: fields[5], Out: fields[6],
					Source: fields[7], Destination: fields[8],
					ExtendedFields: extendedFields})
			}
		}
	}

}
