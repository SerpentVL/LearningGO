package array

import (
	"strings"
)

type Bug struct {
	Type           string
	Category       string
	CWE            int
	TranslatedType string
}

func FindCWE(bug_type string, bug_category string) *Bug {
	// Simple cases
	for _, bug := range bugs {
		if strings.EqualFold(bug_type, bug.Type) && strings.EqualFold(bug_category, bug.Category) {
			return &bug
		}
	}

	// TODO: Special cases
	return nil
}
