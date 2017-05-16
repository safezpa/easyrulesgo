package core

import "github.com/CrowdStrike/easyrulesgo/api"


// ByPriority implements sort.Interface for []api.Rule based on
// the ByPriority field.
type ByPriority []api.Rule

func (r ByPriority) Len() int           { return len(r) }
func (r ByPriority) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByPriority) Less(i, j int) bool { return r[i].Priority() < r[j].Priority() }
