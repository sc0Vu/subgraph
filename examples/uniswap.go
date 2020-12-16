package main

import (
	"context"
	"fmt"
	"github.com/sc0Vu/subgraph/uniswapv2"
)

func main() {
	cli := uniswapv2.NewUniswapV2Client("zzz")
	ethPrice, err := cli.Bundles(context.TODO(), 1, 11411111)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Current ethereum price: %f\n", ethPrice)
}
