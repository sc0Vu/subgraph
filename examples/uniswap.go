package main

import (
	"context"
	"fmt"
	"github.com/sc0Vu/subgraph/uniswapv2"
	"os"
	"strconv"
)

func main() {
	// const bn = 11448502
	const bn = 11455055
	token := os.Getenv("SUBGRAPH_TOKEN")
	pairAddress := os.Getenv("UNISWAP_ADDRESS")
	cli := uniswapv2.NewUniswapV2Client(token)
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

	pair, err := cli.PairsWithBN(context.TODO(), pairAddress, bn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Pair information in %d: %+v\n", bn, pair)
	pairNow, err := cli.Pairs(context.TODO(), "0xb6a0d0406772ac3472dc3d9b7a2ba4ab04286891")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Current pair information: %+v\n", pairNow)
	reserveETHNow, err := strconv.ParseFloat(string(pairNow.TrackedReserveETH), 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	liquidityNow := reserveETHNow * ethPriceNow
	fmt.Printf("Current liquidity: %f\n", liquidityNow)
	volumeNow, err := strconv.ParseFloat(string(pairNow.VolumeUSD), 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	volume, err := strconv.ParseFloat(string(pair.VolumeUSD), 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	oneDayVolume := volumeNow - volume
	fmt.Printf("One day volume: %f\n", oneDayVolume)
	fee := oneDayVolume * 0.003
	fmt.Printf("Fee: %f\n", fee)
}
