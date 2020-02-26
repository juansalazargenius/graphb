package graphb

import (
	"regexp"
	"strings"
)

// checks the validity of a name according to the spec: http://facebook.github.io/graphql/October2016/#sec-Names
var validName = regexp.MustCompile(`^([A-Za-z0-9_]+)(\(\$[A-Za-z0-9_]+:[A-Za-z]+\))?$`)

func isValidOperationType(Type operationType) bool {
	low := strings.ToLower(string(Type))
	return low == "query" || low == "mutation" || low == "subscription"
}

const (
	// syntax tokens
	tokenLB     = "{" // Left Brace
	tokenRB     = "}" // Right Brace
	tokenLP     = "(" // Left Parenthesis
	tokenRP     = ")" // Right Parenthesis
	tokenColumn = ":"
	tokenComma  = ","
	tokenSpace  = " "
)
