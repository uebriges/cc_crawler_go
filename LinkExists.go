package main

import (
	"strings"
)

// ------------------------------------------------
// Helper function: Checks if link already exists
// in slice and returns true/false
// ------------------------------------------------
func LinkExists(mapOfLinks map[string]bool, newLink string) (exists bool) {
	// Ignore some relative links
	// Possible improvement would need to keep track of "link level"
	if strings.Contains(newLink, "../") || strings.HasPrefix(newLink, "page-") || !strings.HasPrefix(newLink, "catalogue/") {
		// fmt.Printf("Links ignored: %s\n", newLink)
		exists = true
		return
	}
	_, exists = mapOfLinks[newLink]
	return
}
