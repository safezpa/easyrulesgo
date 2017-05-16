package main

import (
	"fmt"
	"testing"

	core "github.com/CrowdStrike/easyrulesgo/core"
)


type helloWorldRule struct {
	input string
	core.BasicRule
}

func (h *helloWorldRule) Evaluate() bool {
	return h.input == "yes"
}

func (h *helloWorldRule) Execute() error {
	fmt.Printf("hello world! from [%s]\n", h.Name())
	return nil
}


type helloWorldRule2 struct {
	input string
	core.BasicRule
}

func (h *helloWorldRule2) Evaluate() bool {
	return h.input == "yes"
}

func (h *helloWorldRule2) Execute() error {
	fmt.Printf("hello world2! from [%s]\n", h.Name())
	return nil
}


func TestBasicRule(t *testing.T) {
	fmt.Println("starting rules engine")

	hw := &helloWorldRule{}
	hw.SetName("Hello World Rule")
	hw.SetPriority(1)
	hw.input = "yes"

	hw2 := &helloWorldRule2{}
	hw2.SetName("Hello World Rule2")
	hw2.SetPriority(2)
	hw2.input = "yes"


	re := core.NewDefaultRulesEngine().
	SkipOnFirstFailedRule(true).
	SkipOnFirstAppliedRule(true)

	re.AddRule(hw)
	re.AddRule(hw2)

	err := re.FireRules()
	if err != nil {
		fmt.Println("rules failed!")
	}
}
