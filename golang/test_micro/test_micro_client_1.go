package main

import (
	"log"
	//"time"
	"fmt"
	"strconv"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-nats"
	//as "github.com/aerospike/aerospike-client-go"

	//config "github.com/yougroupteam/you-config"
	//handler "github.com/yougroupteam/you-ledger/ledger"
	pbforex "github.com/yougroupteam/you-pb/go/pb/services/you_forex_cache"
	"github.com/yougroupteam/you-pb/go/pb/money"

	"golang.org/x/net/context"
)


func main() {
	// Parse our command line options
	cmd.Init()

	s := nats. NewService(
		micro.Name("go.micro.client.you-forex-cache"),
		//micro.RegisterTTL(time.Second*30),
		//micro.RegisterInterval(time.Second*10),
	)
	//s.Init()

	fxQuoteClient := pbforex.NewFxQuoteClient("go.micro.srv.you-forex-cache", s.Client())

	/*
	type GetDetailedRateRequest struct {
		SellAmount  *pb_money1.Money    `protobuf:"bytes,1,opt,name=sellAmount" json:"sellAmount,omitempty"`
		BuyCurrency *pb_money1.Currency `protobuf:"bytes,2,opt,name=buyCurrency" json:"buyCurrency,omitempty"`
		FixedSide   string              `protobuf:"bytes,3,opt,name=fixedSide" json:"fixedSide,omitempty"`
	}

	type Money struct {
		CurrencyCode CurrencyCode `protobuf:"varint,1,opt,name=currency_code,json=currencyCode,enum=pb.money.CurrencyCode" json:"currency_code,omitempty"`
		ValueMicros  int64        `protobuf:"varint,2,opt,name=value_micros,json=valueMicros" json:"value_micros,omitempty"`
	}

	*/

	// GetDetailedRate(from, to, fixedSide, amount string) (*pbforex.Rate, error)
	fixedSide := "sell"

	buyCurrencyCode, err := pbmoney.CurrencyCodeSimpleValueOf("USD")
	if err != nil {
		log.Fatal(err)
	}
	sellCurrencyCode, err := pbmoney.CurrencyCodeSimpleValueOf("SGD")
	if err != nil {
		log.Fatal(err)
	}
	buyCurrency, ok := pbmoney.CurrencyCodeToCurrency[buyCurrencyCode]
	if !ok {
		log.Fatal(fmt.Errorf("Could not find currency with code: %d", buyCurrencyCode))
	}
	amountFloat, err := strconv.ParseFloat("1000.0", 64)
	if err != nil {
		log.Fatal(err)
	}
	req := new(pbforex.GetDetailedRateRequest)
	req.SellAmount = pbmoney.NewMoneyFloat(sellCurrencyCode, amountFloat)
	req.BuyCurrency = buyCurrency
	req.FixedSide = fixedSide

	log.Printf("request: %v\n", *req)
	
	resp, err := fxQuoteClient.GetDetailedRate(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("response: %v\n", *resp)
}
