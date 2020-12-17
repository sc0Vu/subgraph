package uniswapv2

import (
	"context"
	"testing"
)

const (
	targetBN = 11111111
	bundleID = 1
	// test pair id for WBTC-ETH
	pairID = "0xbb2b8038a1640196fbe3e38816f3e67cba72d940"
)

func newClient() (uniCli UniswapV2Client) {
	uniCli = NewUniswapV2Client("")
	return
}

func newCtx() (ctx context.Context) {
	ctx = context.Background()
	return
}

// TestBundles test subgraph bundles api
func TestBundles(t *testing.T) {
	c := newClient()
	ctx := newCtx()
	ethPriceNow, err := c.Bundles(ctx, bundleID)
	if err != nil {
		t.Fatal(err)
	}
	if ethPriceNow < 0 {
		t.Fatalf("Eth price should not be zero")
	}
	ethPriceOld, err := c.BundlesWithBN(ctx, bundleID, targetBN)
	if err != nil {
		t.Fatal(err)
	}
	if ethPriceOld < 0 {
		t.Fatalf("Eth price should not be zero")
	}
	if ethPriceNow == ethPriceOld {
		t.Fatalf("Eth price old should not be the same with eth price now")
	}
}

// TestPairs test subgraph pairs api
func TestPairs(t *testing.T) {
	c := newClient()
	ctx := newCtx()
	pairNow, err := c.Pairs(ctx, pairID)
	if err != nil {
		t.Fatal(err)
	}
	if pairNow.ID != pairID {
		t.Fatalf("Pair ID not equal")
	}
	pairOld, err := c.PairsWithBN(ctx, pairID, targetBN)
	if err != nil {
		t.Fatal(err)
	}
	if pairOld.ID != pairID {
		t.Fatalf("Pair ID not equal")
	}
	if pairNow.Token0.Name != pairOld.Token0.Name {
		t.Fatalf("Pairs token0 name not equal")
	}
	if pairNow.Token0.Symbol != pairOld.Token0.Symbol {
		t.Fatalf("Pairs token0 name not equal")
	}
	if pairNow.Token1.Name != pairOld.Token1.Name {
		t.Fatalf("Pairs token1 name not equal")
	}
	if pairNow.Token1.Symbol != pairOld.Token1.Symbol {
		t.Fatalf("Pairs token1 name not equal")
	}
}
