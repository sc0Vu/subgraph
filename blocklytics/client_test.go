package blocklytics

import (
	"context"
	"strconv"
	"testing"
)

const (
	targetBN = 11111111
)

func newClient() (uniCli BlocklyticsClient) {
	uniCli = NewBlocklyticsClient("")
	return
}

func newCtx() (ctx context.Context) {
	ctx = context.Background()
	return
}

// TestBlocks test subgraph blocks api
func TestBlocks(t *testing.T) {
	c := newClient()
	ctx := newCtx()
	blocks, err := c.Blocks(ctx, 10, 0, targetBN, "number desc")
	if err != nil {
		t.Fatal(err)
	}
	if len(blocks) != 10 {
		t.Fatalf("Should fetch 10 blocks")
	}
	for _, block := range blocks {
		number, err := strconv.ParseInt(string(block.Number), 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		if number < targetBN {
			t.Fatalf("Should fetch blocks that block number greater than %d", targetBN)
		}
	}
}
