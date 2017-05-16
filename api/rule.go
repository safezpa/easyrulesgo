package api


// Rule is the basic interface that defines what a business Rule should have
type Rule interface {
	Name() string
	Description() string
	Priority() int
	Evaluate() bool
	Execute() error
}