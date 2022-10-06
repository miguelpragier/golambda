package apigwhandler

import "strings"

const PathParameterIdentifier = ":"

// RouteMatches compares the given rawPath with a route, and returns true if they match their static parts
// It doesn't accept partial route matching
func RouteMatches(route, rawPath string) bool {
	sliceRoute := strings.Split(route, "/")
	slicePath := strings.Split(rawPath, "/")

	if len(sliceRoute) != len(slicePath) {
		return false
	}

	for i, routePiece := range sliceRoute {
		if strings.Contains(routePiece, PathParameterIdentifier) {
			continue
		}

		if routePiece != slicePath[i] {
			return false
		}
	}

	return true
}
