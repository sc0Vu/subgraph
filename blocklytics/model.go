package blocklytics

import (
	"github.com/shurcooL/graphql"
)

// Block represents graphql model of block
type Block struct {
	ID         graphql.String
	Number     graphql.String
	Timestamp  graphql.String
	Author     graphql.String
	Difficulty graphql.String
	GasUsed    graphql.String
	GasLimit   graphql.String
}
