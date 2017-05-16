package core

import (
	"github.com/CrowdStrike/easyrulesgo/api"
)

// CompositeRule is a rule type that requires all the rules to evaluate to true to execute any of them
type CompositeRule struct {
	BasicRule
	rules []api.Rule
}

// Evaluate ensure that all the rules evaluate to true, otherwise returns false
func (c *CompositeRule) Evaluate() bool {
	if len(c.rules) == 0 {
		return false
	}

	for _, rule := range c.rules {
		if !rule.Evaluate() {
			return false
		}
	}
	return true
}

// Execute all the rules passed, run each rules execute method
func (c *CompositeRule) Execute() error {

	for _, rule := range c.rules {
		err := rule.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

// AddRule add a rule to the composite
func (c *CompositeRule) AddRule(rule api.Rule) {
	c.rules = append(c.rules, rule)
}

// GetName returns the name of the rule
func (c *CompositeRule) GetName() string {
	return "Composite Rule"
}

// GetDescription returns a brief description of the rule
func (c *CompositeRule) GetDescription() string {
	return "Composite Rule ensures that all the rules included match"
}