package cmd

type ChainRule struct {
	Destination    string
	ExtendedFields string
	In             string
	Out            string
	Prot           string
	Source         string
	Target         string
}

type IPChain struct {
	Name  string
	Rules []ChainRule
}

func (chain *IPChain) AddRule(rule ChainRule) []ChainRule {
	chain.Rules = append(chain.Rules, rule)
	return chain.Rules
}
