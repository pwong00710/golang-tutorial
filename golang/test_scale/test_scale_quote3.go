package main

import (
	"fmt"
	"log"

	"time"

	config "github.com/yougroupteam/you-config"
	//"github.com/yougroupteam/you-forex-cache/cache"
	scalehandler "github.com/yougroupteam/you-forex-cache/handlers/scale"
)

func main() {
	from := "SGD"
	to := "USD"
	fixedSide := "Sell"

	cnf := config.NewConfig(nil, true, false)

	handler, err := scalehandler.New(cnf)
	if err != nil {
		log.Fatal(err)
	}

	//fxProviderCache := cache.New(handler.GetName())

	for true {
		err = handler.ExchangeRate(nil, "nil", from, to, fixedSide)
		if err != nil {
			log.Print(err)
		}
		time.Sleep(5 * time.Second)
		//fmt.Printf("test:main:%v %v %v\n", fxProviderCache.FXProvider, fxProviderCache.ExchangeRates, fxProviderCache.LastUpdate)
	}
	fmt.Printf("Exit!")
}
