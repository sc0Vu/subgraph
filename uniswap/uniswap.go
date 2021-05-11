package uniswap

import (
	"context"
	"strconv"

	"github.com/sc0Vu/graphql"
)

// TODO: remove reflect usage?
// TODO: big decimal type in case of precision?

// UniswapClient represents uniswapv2 subgraph client
type UniswapClient struct {
	c     *graphql.Client
	token string
}

// NewUniswapClient returns uniswap subgraph client
func NewUniswapClient(token string) (uniCli UniswapClient) {
	uniCli.c = graphql.NewClient(v1Endpoint, nil)
	uniCli.token = token
	return
}

// BundlesWithBN returns the price of eth in given block number
func (uniCli *UniswapClient) BundlesWithBN(ctx context.Context, id, bn int) (ethPrice float64, err error) {
	var query struct {
		Bundle Bundle `graphql:"bundle(id: $id, block:{number: $bn})"`
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

// Bundles returns the current price of eth
func (uniCli *UniswapClient) Bundles(ctx context.Context, id int) (ethPrice float64, err error) {
	var query struct {
		Bundle Bundle `graphql:"bundle(id: $id)"`
	}
	variables := map[string]interface{}{
		"id": graphql.Int(id),
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

// TokensWithBN returns the token in uniswap of given address and block number
func (uniCli *UniswapClient) TokensWithBN(ctx context.Context, id string, bn int) (token Token, err error) {
	var query struct {
		Token Token `graphql:"token(id: $id, where:{id: $tokenID}, block:{number: $bn})"`
	}
	variables := map[string]interface{}{
		"id":      graphql.String(id),
		"tokenID": graphql.String(id),
		"bn":      graphql.Int(bn),
	}
	err = uniCli.c.Query(ctx, &query, variables)
	if err != nil {
		return
	}
	token = query.Token
	return
}

// Tokens returns the token in uniswap of given address
func (uniCli *UniswapClient) Tokens(ctx context.Context, id string) (token Token, err error) {
	var query struct {
		Token Token `graphql:"token(id: $id, where:{id: $id})"`
	}
	variables := map[string]interface{}{
		"id": graphql.String(id),
	}
	err = uniCli.c.Query(ctx, &query, variables)
	if err != nil {
		return
	}
	token = query.Token
	return
}
