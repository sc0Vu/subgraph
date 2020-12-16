package uniswapv2

import (
	"context"
	"github.com/shurcooL/graphql"
	"strconv"
)

const endpoint = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"

// TODO: add token
// TODO: remove reflect usage?
// TODO: big decima type in case of precision?

// UniswapV2Client represents uniswapv2 subgraph client
type UniswapV2Client struct {
	c     *graphql.Client
	token string
}

// NewUniswapV2Client returns uniswapv2 subgraph client
func NewUniswapV2Client(token string) (uniCli UniswapV2Client) {
	uniCli.c = graphql.NewClient(endpoint, nil)
	uniCli.token = token
	return
}

// Bundles returns the price of eth
func (uniCli *UniswapV2Client) Bundles(ctx context.Context, id, bn int) (ethPrice float64, err error) {
	var query struct {
		Bundle struct {
			ID       graphql.ID
			EthPrice graphql.String
		} `graphql:"bundle(id: $id, block:{number: $bn})"`
	}
	variables := map[string]interface{}{
		"id": graphql.Int(id),
		"bn": graphql.Int(bn),
	}
	err = uniCli.c.Query(ctx, &query, variables)
	if err != nil {
		return
	}
	if ethPrice, err = strconv.ParseFloat(string(query.Bundle.EthPrice), 64); err != nil {
		return
	}
	return
}
