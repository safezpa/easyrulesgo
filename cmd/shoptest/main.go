package main

import (
	"fmt"

	core "github.com/CrowdStrike/easyrulesgo/core"
)

// example taken from http://www.easyrules.org/tutorials/shop-tutorial.html

type person struct {
	name  string
	age   int
	adult bool
}


/*-------- AGE RULE --------*/
type ageRule struct {
	core.BasicRule
	person *person
}

func (a *ageRule) Evaluate() bool {
	return a.person.age > a.adultAge()
}

func (a *ageRule) Execute() error {
	a.person.adult = true
	fmt.Printf("person %s has been marked as adult", a.person.name)
	return nil
}

func (a *ageRule) adultAge() int {
	return 21
}

func newAgeRule(p *person) *ageRule {
	ar := &ageRule{}
	ar.person = p
	return ar
}


/*-------- ALCOHOL RULE --------*/

type alcoholRule struct {
	core.BasicRule
	person *person
}

func (a *alcoholRule) Evaluate() bool {
	return !a.person.adult
}

func (a *alcoholRule) Execute() error {
	fmt.Printf("Shop: Sorry %s you are not allowed to buy alcohol", a.person.name)
	return nil
}

func newAlcoholRule(p *person) *alcoholRule {
	ar := &alcoholRule{}
	ar.person = p
	return ar
}

func main() {

	p := &person{name: "Tom", age: 14}
	ar := newAgeRule(p)
	alcr := newAlcoholRule(p)

	re := core.NewDefaultRulesEngine()
	re.AddRule(ar)
	re.AddRule(alcr)

	err := re.FireRules()
	if err != nil {
		fmt.Println("rules failed!")
	}
}


