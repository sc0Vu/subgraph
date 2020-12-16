package main

import (
	"context"
	"fmt"
	"github.com/sc0Vu/subgraph/uniswapv2"
)

func main() {
	const bn = 11411111
	cli := uniswapv2.NewUniswapV2Client("zzz")
	ethPrice, err := cli.BundlesWithBN(context.TODO(), 1, bn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Ethereum price in %d: %f\n", bn, ethPrice)
	ethPriceNow, err := cli.Bundles(context.TODO(), 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Current ethereum price: %f\n", ethPriceNow)
}
