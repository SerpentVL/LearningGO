package array

import (
	"strings"
)

type Bug struct {
	Priority       string
	Type           string
	EndType        string
	Category       string
	CWE            int
	TranslatedType string
}

func FindCWE(bug_type string, bug_category string) *Bug {
	// Simple cases
	for _, bug := range Bugs {
		if strings.EqualFold(bug_category, bug.Category) {
			if strings.EqualFold(bug_type, bug.Type) ||
				strings.HasPrefix(bug_type, bug.Type) && strings.HasSuffix(bug_type, bug.EndType) { // complicated type of bug
				return &bug
			}
		}
	}

	// TODO: Special cases
	return nil
}
