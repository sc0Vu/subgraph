package uniswapv2

import (
	"github.com/sc0Vu/graphql"
)

// Bundle represent graphql model of Bundle
type Bundle struct {
	ID       graphql.ID
	EthPrice graphql.String
}

// Token represent graph model of token
type Token struct {
	ID                 graphql.ID
	Symbol             graphql.String
	Name               graphql.String
	Decimals           graphql.String
	TotalSupply        graphql.String
	TradeVolume        graphql.String
	TradeVolumeUSD     graphql.String `graphql:"tradeVolumeUSD"`
	UntrackedVolumeUSD graphql.String `graphql:"untrackedVolumeUSD"`
	TXCount            graphql.String
	TotalLiquidity     graphql.String
	DerivedETH         graphql.String `graphql:"derivedETH"`
}

// Pair represent graphql model of Pair
type Pair struct {
	ID                     graphql.ID
	Token0                 Token
	Token1                 Token
	TrackedReserveETH      graphql.String `graphql:"trackedReserveETH"`
	VolumeUSD              graphql.String `graphql:"volumeUSD"`
	UntrackedVolumeUSD     graphql.String `graphql:"untrackedVolumeUSD"`
	TXCount                graphql.String `graphql:"txCount"`
	CreatedAtTimestamp     graphql.String `graphql:"createdAtTimestamp"`
	CreatedAtBlockNumber   graphql.String `graphql:"createdAtBlockNumber"`
	LiquidityProviderCount graphql.String
	Reserve0               graphql.String
	Reserve1               graphql.String
}
