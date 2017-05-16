package core


const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)


// BasicRule represents an embeddable type that can be used with your actual rules to provide some of the basic helper
// functionality
type BasicRule struct {
	name string
	description string
	priority int
}

// SetName set the name of the rule
func (b *BasicRule) SetName(name string) {
	b.name = name
}

// Name returns the name of the rule
func (b *BasicRule) Name() string {
	return b.name
}

// Description returns the description of the rule
func (b *BasicRule) Description() string {
	return b.description
}

// SetDescription set the description of the rule
func (b *BasicRule) SetDescription(desc string)  {
	b.description = desc
}


// Priority returns the current priority, note that a priority of 0 or less defaults to the lowest priority. Using 1-n
func (b *BasicRule) Priority() int {
	if b.priority <= 0 {
		return maxInt
	}
	return b.priority
}

// SetPriority sets the priority of your rule, defaults to MAXINT as the lowest priority
func (b *BasicRule) SetPriority(priority int) {
	b.priority = priority
}




