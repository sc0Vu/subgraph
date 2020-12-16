package uniswapv2

import (
	"github.com/shurcooL/graphql"
)

// Bundle represent graphql model of Bundle
type Bundle struct {
	ID       graphql.ID
	EthPrice graphql.String
}
