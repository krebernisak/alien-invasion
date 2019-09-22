package types

import (
	"fmt"
	"sort"
	"strings"
)

// Link type representing connection between Nodes
type Link struct {
	Key   string
	Nodes []string
}

// NewLink creates a Link with sorted key
func NewLink(nodes ...string) Link {
	sort.Strings(nodes)
	key := strings.Join(nodes, "_")
	return Link{key, nodes}
}

// String representation of a Link
func (l Link) String() string {
	return fmt.Sprintf("key=%s nodes=%s\n", l.Key, l.Nodes)
}
