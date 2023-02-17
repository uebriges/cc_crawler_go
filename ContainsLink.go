package main

import (
	"strings"
)

// ------------------------------------------------
// Helper function: Checks if link already exists
// in map and returns true/false
// ------------------------------------------------
func ContainsLink(mapOfLinks map[string]bool, newLink string) (containsLink bool) {
	// Ignore some relative links
	// Possible improvement would need to keep track of "link level"
	if strings.Contains(newLink, "../") || strings.HasPrefix(newLink, "page-") || !strings.HasPrefix(newLink, "catalogue/") {
		// fmt.Printf("Links ignored: %s\n", newLink)
		containsLink = true
		return
	}
	_, containsLink = mapOfLinks[newLink]
	return
}
