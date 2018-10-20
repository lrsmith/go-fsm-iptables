package cmd

import "fmt"

func FollowChain(chain string) {
	start := IPChains[chain]
	for order, rule := range start.Rules {

		label := fmt.Sprintf("%d: ( %s,%s,%s,%s,%s : %s )", order, rule.Prot, rule.In, rule.Out, rule.Source, rule.Destination, rule.ExtendedFields)
		switch rule.Target {
		case "ACCEPT", "DROP", "REJECT":
			PrintNode(chain, rule.Target, label)
		case "MARK", "RETURN":
			PrintNode(chain, rule.Target, label)
			PrintNode(rule.Target, chain, "")
		default:
			PrintNode(chain, rule.Target, label)
			FollowChain(rule.Target)
		}
	}
}
