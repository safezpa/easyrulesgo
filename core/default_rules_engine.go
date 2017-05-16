package core

import (
	api "github.com/CrowdStrike/easyrulesgo/api"
	"fmt"
	"sort"
)

// DefaultRulesEngine is a basic rules engine implementation that does an evaluate/execute pattern on a set of given rules
type DefaultRulesEngine struct {
	rules []api.Rule
	skipOnFirstFailedRule bool
	skipOnFirstAppliedRule bool
	logger api.Logger
}

// AddRule adds a rule for eval/execute
func (d *DefaultRulesEngine) AddRule(rule api.Rule) {
	d.rules = append(d.rules, rule)
}

// GetRules returns the rules currently registered
func (d *DefaultRulesEngine) GetRules() []api.Rule {
	return 	d.rules
}

// ClearRules deletes all the rules currently in place
func (d *DefaultRulesEngine) ClearRules() {
	d.rules = []api.Rule{}
}

// FireRules runs the eval/execute loop on your rules while taking prioritization into account
func (d *DefaultRulesEngine) FireRules() error {

	// todo need to sort rules here based on priority
	d.sortRules()
	err := d.applyRules()
	if err != nil {
		return err
	}
	return nil
}


// SkipOnFirstFailedRule if a rule returns an error when executing, stop processing immediately
func (d *DefaultRulesEngine) SkipOnFirstFailedRule(skip bool) *DefaultRulesEngine {
	d.skipOnFirstFailedRule = skip
	return d
}

// SkipOnFirstAppliedRule will stop evaluation once the first rule is applied
func (d *DefaultRulesEngine) SkipOnFirstAppliedRule(skip bool) *DefaultRulesEngine {
	d.skipOnFirstAppliedRule = skip
	return d
}

// SetLogger allows you to pass your own logger instance in to incorporate logging into your app. Highly encouraged to use your own here
// just implement the interface from api/logger.go
func (d *DefaultRulesEngine) SetLogger(logger api.Logger) *DefaultRulesEngine {
	d.logger = logger
	return d
}

func (d *DefaultRulesEngine) sortRules() {
	sort.Sort(ByPriority(d.rules))
}

func (d *DefaultRulesEngine) applyRules() error {
	for _, rule := range d.rules {
		if rule.Evaluate() {
			err := rule.Execute()
			if err != nil {
				if d.skipOnFirstFailedRule {
					d.logger.Log(WARN, fmt.Errorf("rule [%s] failed with error [%s]", rule.Name(), err))
				}
				return fmt.Errorf("rule [%s] failed to execute; %s", rule.Name(), err)
			}
			if d.skipOnFirstAppliedRule {
				d.logger.Log(INFO, fmt.Errorf("rule [%s] executed, skipping further processing as requested", rule.Name()))
				return nil
			}
		}
	}
	return nil
}


// NewDefaultRulesEngine returns a new Rules Engine instance
func NewDefaultRulesEngine() *DefaultRulesEngine {
	d := &DefaultRulesEngine{}
	d.rules = []api.Rule{}
	d.logger = &DefaultLogger{}
	return d
}