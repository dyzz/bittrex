package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gocarina/gocsv"
	"github.com/toorop/go-bittrex"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

var _ = spew.Dump

const (
	Key    = "92720a41e5574f1fb20b9e16e6579581"
	Secret = "c94fd8a2ade5431e80aae14697c29382"
)

var api = bittrex.New(Key, Secret)

func listMarkets() {
	markets, err := api.GetMarkets()
	if err != nil {
		log.Println(err)
	}
	marketNames := make([]string, len(markets))
	for i, market := range markets {
		marketNames[i] = market.MarketName
	}
	fmt.Println(marketNames)
}

func getCandle(market string, interval string) {
	candles, err := api.GetHisCandles(market, interval, time.Now().Unix()*1000)
	if err != nil {
		log.Println(err)
	}

	fileName := fmt.Sprintf("%s-%s-%s.csv", market, interval, time.Now().Format("2006-01-02-15-04"))
	outputFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	if err := gocsv.MarshalFile(&candles, outputFile); err != nil {
		log.Fatal(err)
	}
	log.Println("done")
}

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "list all market symbols",
			Action: func(c *cli.Context) error {
				listMarkets()
				return nil
			},
		},
		{
			Name:  "candle",
			Usage: "get candles",
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 2 {
					fmt.Println("./bittrex candle market interval(oneMin, fiveMin, thirtyMin, hour, day)")
					return nil
				}
				getCandle(c.Args()[0], c.Args()[1])
				return nil
			},
		},
	}

	app.Run(os.Args)
}
