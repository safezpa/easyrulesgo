package easyrulesgo

import (
	"testing"
	"os"
	"github.com/CrowdStrike/gomock/gomock"
	"github.com/CrowdStrike/easyrulesgo/api"
	"github.com/CrowdStrike/easyrulesgo/core"
	"fmt"
)


func TestMain(m *testing.M) {
	os.Exit(m.Run())
}


func TestBasicRule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRule1 := api.NewMockRule(ctrl)
	mockRule1.EXPECT().Evaluate().Return(true)
	mockRule1.EXPECT().Execute().Return(nil)
	re := core.NewDefaultRulesEngine()
	re.AddRule(mockRule1)
	re.FireRules()
}

// verify that if we fail the first rule that the 2nd rule never gets called
func TestFirstFailSkip(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRule1 := api.NewMockRule(ctrl)
	mockRule1.EXPECT().Evaluate().Return(true)
	mockRule1.EXPECT().Execute().Return(fmt.Errorf("unknown error occured during execution"))
	mockRule1.EXPECT().Priority().Return(1)
	mockRule1.EXPECT().Name().Return("Hello World rule")

	// shouldn't even hit this rule, just call it's get Priority
	mockRule2 := api.NewMockRule(ctrl)
	mockRule2.EXPECT().Priority().Return(2)
	mockRule1.EXPECT().Name().Return("Hello World2 rule")


	re := core.NewDefaultRulesEngine()
	re.SkipOnFirstFailedRule(true)
	re.AddRule(mockRule1)
	re.AddRule(mockRule2)
	re.FireRules()
}

// verify that if we match on the first rule that the second rule never gets run
func TestFirstAppliedSkip(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRule1 := api.NewMockRule(ctrl)
	mockRule1.EXPECT().Evaluate().Return(true)
	mockRule1.EXPECT().Execute().Return(nil)
	mockRule1.EXPECT().Priority().Return(1)
	mockRule1.EXPECT().Name().Return("Hello World rule")

	// shouldn't even hit this rule, just call it's get Priority
	mockRule2 := api.NewMockRule(ctrl)
	mockRule2.EXPECT().Priority().Return(2)


	re := core.NewDefaultRulesEngine()
	re.SkipOnFirstAppliedRule(true)
	re.AddRule(mockRule1)
	re.AddRule(mockRule2)
	re.FireRules()
}

// verify that execute will  not be called if any rules failed
func TestCompositeRuleFailsOnFalse(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRule1 := api.NewMockRule(ctrl)
	mockRule1.EXPECT().Evaluate().Return(true)


	mockRule2 := api.NewMockRule(ctrl)
	mockRule2.EXPECT().Evaluate().Return(false)

	compositeRule := &core.CompositeRule{}
	compositeRule.AddRule(mockRule1)
	compositeRule.AddRule(mockRule2)

	re := core.NewDefaultRulesEngine()
	re.AddRule(compositeRule)
	re.FireRules()
}


// verify that execute will be called if the rules return true evaluations
func TestCompositeRuleExecutesOnTrue(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRule1 := api.NewMockRule(ctrl)
	mockRule1.EXPECT().Evaluate().Return(true)
	mockRule1.EXPECT().Execute().Return(nil)


	mockRule2 := api.NewMockRule(ctrl)
	mockRule2.EXPECT().Evaluate().Return(true)
	mockRule2.EXPECT().Execute().Return(nil)

	compositeRule := &core.CompositeRule{}
	compositeRule.AddRule(mockRule1)
	compositeRule.AddRule(mockRule2)

	re := core.NewDefaultRulesEngine()
	re.AddRule(compositeRule)
	re.FireRules()
}