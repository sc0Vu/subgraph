package uniswap

import (
	"context"
	// "strconv"

	"github.com/sc0Vu/graphql"
)

// TODO: remove reflect usage?
// TODO: big decimal type in case of precision?

// UniswapV2Client represents uniswapv2 subgraph client
type UniswapV2Client struct {
	UniswapClient
}

// NewUniswapV2Client returns uniswapv2 subgraph client
func NewUniswapV2Client(token string) (uniCli UniswapV2Client) {
	uniCli.c = graphql.NewClient(v2Endpoint, nil)
	uniCli.token = token
	return
}

// PairsWithBN returns the pair in uniswap of given address and block number
func (uniCli *UniswapV2Client) PairsWithBN(ctx context.Context, id string, bn int) (pair Pair, err error) {
	var query struct {
		Pair Pair `graphql:"pair(id: $id, where:{id: $pairID}, block:{number: $bn})"`
	}
	variables := map[string]interface{}{
		"id":     graphql.String(id),
		"pairID": graphql.String(id),
		"bn":     graphql.Int(bn),
	}
	err = uniCli.c.Query(ctx, &query, variables)
	if err != nil {
		return
	}
	pair = query.Pair
	return
}

// Pairs returns the pair in uniswap of given address
func (uniCli *UniswapV2Client) Pairs(ctx context.Context, id string) (pair Pair, err error) {
	var query struct {
		Pair Pair `graphql:"pair(id: $id, where:{id: $id})"`
	}
	variables := map[string]interface{}{
		"id": graphql.String(id),
	}
	err = uniCli.c.Query(ctx, &query, variables)
	if err != nil {
		return
	}
	pair = query.Pair
	return
}
