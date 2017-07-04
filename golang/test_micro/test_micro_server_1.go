package main

import (
	"log"
	"time"

	"golang.org/x/net/context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-nats"

	//config "github.com/yougroupteam/you-config"
	"github.com/yougroupteam/you-pb/go/pb/money"
	pbforex "github.com/yougroupteam/you-pb/go/pb/services/you_forex_cache"
)

type Handler struct {
	//cnf *config.Config
}

//	GetDetailedRate(context.Context, *GetDetailedRateRequest, *Rate) error
//	RetrieveMultipleRates(context.Context, *RetrieveMultipleRatesRequest, *MultiRate) error

func (h *Handler) GetDetailedRate(ctx context.Context, in *pbforex.GetDetailedRateRequest, out *pbforex.Rate) error {
	log.Printf("request: %v\n", *in)
	/*
		rate := new(pbforex.Rate)
		rate.Rate = new(pbmoney.ExchangeRate)
		var err error
		rate.Rate.From, err = pbmoney.CurrencyCodeSimpleValueOf("SGD")
		if err != nil {
			log.Fatal(err)
		}
		rate.Rate.To, err = pbmoney.CurrencyCodeSimpleValueOf("USD")
		if err != nil {
			log.Fatal(err)
		}
		rate.Rate.ValueMicros = int64(1.2992 * 1000000.0)
		out = rate
	*/

	out.Rate = new(pbmoney.ExchangeRate)
	var err error

	out.Rate.From, err = pbmoney.CurrencyCodeSimpleValueOf("SGD")
	if err != nil {
		log.Fatal(err)
	}
	out.Rate.To, err = pbmoney.CurrencyCodeSimpleValueOf("USD")
	if err != nil {
		log.Fatal(err)
	}
	out.Rate.ValueMicros = int64(1.2992 * 1000000.0)
	log.Printf("response: %v\n", *out)

	return nil
}

func (h *Handler) RetrieveMultipleRates(ctx context.Context, in *pbforex.RetrieveMultipleRatesRequest, out *pbforex.MultiRate) error {
	log.Printf("request: %v\n", *in)
	multirate := new(pbforex.MultiRate)
	multirate.Rates = []*pbmoney.ExchangeRate{}

	rate := new(pbmoney.ExchangeRate)
	var err error
	rate.From, err = pbmoney.CurrencyCodeSimpleValueOf("SGD")
	if err != nil {
		log.Fatal(err)
	}
	rate.To, err = pbmoney.CurrencyCodeSimpleValueOf("USD")
	if err != nil {
		log.Fatal(err)
	}
	rate.ValueMicros = int64(1.2992 * 1000000.0)
	multirate.Rates = append(multirate.Rates, rate)
	out = multirate
	log.Printf("response: %v\n", *out)

	return nil
}

func (h *Handler) Hello(ctx context.Context, req *pbforex.HelloRequest, rsp *pbforex.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func (h *Handler) GetDetailedRate2(ctx context.Context, in *pbforex.GetDetailedRateRequest, out *pbforex.HelloResponse) error {
	out.Greeting = "Hello World!"
	return nil
}

func main() {
	// Parse our command line options
	cmd.Init()
	//cnf := config.NewConfig(nil, true, true)

	s := nats.NewService(
		micro.Name("go.micro.srv.you-forex-cache"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	s.Init()

	//handler := Handler{cnf: cnf}
	handler := Handler{}
	pbforex.RegisterFxQuoteHandler(s.Server(), &handler)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
