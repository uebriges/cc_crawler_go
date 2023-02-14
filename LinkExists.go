package main

import (
	"strings"
)

// ------------------------------------------------
// Helper function: Checks if link already exists
// in slice and returns true/false
// ------------------------------------------------
func LinkExists(mapOfLinks map[string]bool, newLink string) (exists bool) {
	// Needed to prevent wrong links.
	// Extension would need to keep track of "link level"
	if strings.Contains(newLink, "../") || strings.HasPrefix(newLink, "page-") || !strings.HasPrefix(newLink, "catalogue/") {
		exists = true
		return
	}
	for k := range mapOfLinks {
		if k == newLink {
			exists = true
		} else {
			exists = false
		}
	}
	return
}
