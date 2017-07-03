package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	//"net/url"
	//"strings"
	"time"

	"github.com/pborman/uuid"
	//config "github.com/yougroupteam/you-config"
	//"github.com/yougroupteam/you-forex-cache/cache"
)

var (
	//certFile = flag.String("cert", "someCertFile", "A PEM eoncoded certificate file.")
	//keyFile  = flag.String("key", "someKeyFile", "A PEM encoded private key file.")
	//caFile   = flag.String("CA", "someCertCAFile", "A PEM eoncoded CA's certificate file.")
	//curlCmd = flag.Bool("curlCmd", false, "Curl command")
	currenyPair = flag.String("currencyPair", "SGD-USD", "Currency Pair")
)

type GetQuotesResponse struct {
	ClientID        string `json:"clientID"`
	ClientRequestID string `json:"clientRequestID"`
	Status          string `json:"status"`
	Rates           []Rate `json:"rates"`
}

type Rate struct {
	RateType       string  `json:"rateType"`
	RateQuoteID    string  `json:"rateQuoteID"`
	RateCategoryID string  `json:"rateCategoryID"`
	ValidFrom      int64   `json:"validFrom"`
	ValidTill      int64   `json:"validTill"`
	BuyCurrency    string  `json:"buyCurrency"`
	SellCurrency   string  `json:"sellCurrency"`
	Tenor          int     `json:"tenor"`
	Rate           float64 `json:"rate,string"`
	ClientRateType string  `json:"clientRateType"`
	Status         string  `json:"status"`
}

type PostTransactionsRequest struct {
	MsgUID           string            `json:"msgUid"`
	PostTransactions []PostTransaction `json:"transactions"`
}

type PostTransaction struct {
	ClientID            string  `json:"clientID"`
	ClientRequestID     string  `json:"clientRequestID"`
	ClientRequestTime   int64   `json:"clientRequestTime"`
	RateQuoteID         string  `json:"rateQuoteID"`
	BookingAccountAlias string  `json:"bookingAccountAlias"`
	BuyAmount           float64 `json:"buyAmount,string"`
	SellAmount          float64 `json:"sellAmount,string"`
	ClientRef1          string  `json:"clientRef1"`
	ClientRef2          string  `json:"clientRef2"`
	ClientRef3          string  `json:"clientRef3"`
}

type PostTransactionsResponse struct {
	MsgUID              string                `json:"msgUid"`
	TransactionResponse []TransactionResponse `json:"transactionResponses"`
}

type TransactionResponse struct {
	Status                string  `json:"status"`
	ClientID              string  `json:"clientID"`
	ClientRequestID       string  `json:"clientRequestID"`
	ScaleTxnID            string  `json:"scaleTxnID"`
	ScaleTxnTime          int64   `json:"scaleTxnTime"`
	RateQuoteID           string  `json:"rateQuoteID"`
	BuyCurrency           string  `json:"buyCurrency"`
	SellCurrency          string  `json:"sellCurrency"`
	BuyAmount             float64 `json:"buyAmount,string"`
	SellAmount            float64 `json:"sellAmount,string"`
	Rate                  float64 `json:"rate,string"`
	TradeDate             string  `json:"tradeDate"`
	ValueDate             string  `json:"valueDate"`
	ClientProfits         float64 `json:"clientProfits,string"`
	ClientProfitsCurrency string  `json:"clientProfitsCurrency"`
	ClientRef1            string  `json:"clientRef1"`
	ClientRef2            string  `json:"clientRef2"`
	ClientRef3            string  `json:"clientRef3"`
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	//cnf := config.NewConfig(nil, true, true)
	//fmt.Printf("Values: %v %v %v %v\n", cnf.Scale.CertFile, cnf.Scale.KeyFile, cnf.Scale.CACert, cnf.Scale.InsecureSkipVerify)

	flag.Parse()

	pem := "./GEU00429-1364209.pem"

	// Load client cert
	//cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	cert, err := tls.LoadX509KeyPair(pem, pem)
	//cert, err := tls.LoadX509KeyPair(cnf.Scale.CertFile, cnf.Scale.KeyFile)
	//cert, err := tls.X509KeyPair([]byte(cnf.Scale.CertFile), []byte(cnf.Scale.KeyFile))
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	//caCert, err := ioutil.ReadFile(*caFile)
	//caCert, err := ioutil.ReadFile("./GEU00429-1364209.pem")
	//caCert, err := ioutil.ReadFile(cnf.Scale.CACert)
	caCert, err := ioutil.ReadFile(pem)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}
	//tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Get-Quotes
	fmt.Printf("**********Get Quotes**********\n\n")

	req, err := http.NewRequest("GET", "https://www.demo.fx-scale.com/scale/v1/quotes-service/get-quotes/EZL/*", nil)
	if err != nil {
		log.Fatal(err)
	}

	currentTime := time.Now()

	service := "quotes-service"
	subService := "get-quotes"
	clientID := "EZL"
	transactionID := uuid.New() //"e0d63ba997354229a8e458e6b04d1ced"
	scaleTS := fmt.Sprintf("%v", currentTime.UnixNano())
	apiKey := "a1fc360d-fe95-4182-a9f3-a7180887771a"
	scaleHMAC := genHMAC(service, subService, clientID, scaleTS, transactionID, apiKey)

	req.Header.Add("SCALE_TS", scaleTS)
	req.Header.Add("SCALE_TX_ID", transactionID)
	req.Header.Add("SCALE_HMAC", scaleHMAC)
	req.Header.Add("Content-Type", "application/json")

	//fmt.Printf("TS: %v\n", req.Header.Get("SCALE_TS"))
	//fmt.Printf("TX: %v\n", req.Header.Get("SCALE_TX"))
	//fmt.Printf("HMAC: %v\n", req.Header.Get("SCALE_HMAC"))

	//if *curlCmd == true {
	//	cmdline := "curl -k --header \"SCALE_TS: %v\" --header \"SCALE_TX_ID: %v\" --header \"SCALE_HMAC: %v\" --header \"Content-Type: application/json\" --cert \"%v:1234\" --tlsv1 \"https://www.demo.fx-scale.com/scale/v1/quotes-service/get-quotes/EZL/*\"\n"
	//	fmt.Printf(cmdline, scaleTS, transactionID, scaleHMAC, pem)
	//	return
	//}

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Dump response
	/*
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))

		var quotes GetQuotesResponse

		json.Unmarshal(data, &quotes)

		fmt.Printf("%v\n", quotes)
	*/

	var quotes GetQuotesResponse
	json.NewDecoder(resp.Body).Decode(&quotes)
	fmt.Printf("%v\n", quotes)

	rateQuoteID := "EZL1IJSC8FXK5KB6"

	for i := range quotes.Rates {
		rate := quotes.Rates[i]
		key := fmt.Sprintf("%v-%v", rate.SellCurrency, rate.BuyCurrency)
		//fmt.Printf("%v %v %v ", key, fxCache.ExchangeRates[key], rate.Rate)
		if key == *currenyPair {
			rateQuoteID = rate.RateQuoteID
		}
	}

	/*
		var fxCache *cache.FXProviderCache
		fxCache = cache.New("Scale")

		for i := range quotes.Rates {
			rate := quotes.Rates[i]
			key := fmt.Sprintf("%v%v", rate.SellCurrency, rate.BuyCurrency)
			//fRate, err := strconv.ParseFloat(rate.Rate, 64)
			if err != nil {
				log.Fatal(err)
			}
			fxCache.ExchangeRates[key] = floatUnitsToMicros(rate.Rate)
			//fmt.Printf("%v %v %v ", key, fxCache.ExchangeRates[key], rate.Rate)
		}
		*fxCache.LastUpdate = time.Now()

		fmt.Printf("%v %v\n", fxCache, fxCache.LastUpdate)
	*/

	// Post-Transaction
	fmt.Printf("*******Post Transaction*******\n\n")

	/*
		apiURL := "https://www.demo.fx-scale.com"
		resource := fmt.Sprintf("/scale/v1/transactions-service/post-transactions/%v", clientID)

		u, _ := url.ParseRequestURI(apiURL)
		u.Path = resource
		urlStr := u.String() // "https://api.com/user/"
	*/

	bookingAccountAlias := "SCALE-EZLSG"
	sellAmount := float64(1000)

	txnReq := PostTransactionsRequest{
		MsgUID: uuid.New(),
		PostTransactions: []PostTransaction{
			PostTransaction{
				ClientID:            clientID,
				ClientRequestID:     uuid.New(),
				ClientRequestTime:   time.Now().UnixNano(),
				RateQuoteID:         rateQuoteID,
				BookingAccountAlias: bookingAccountAlias,
				SellAmount:          sellAmount,
			},
		},
	}

	//postData, err := json.Marshal(txnReq)

	postData := new(bytes.Buffer)
	json.NewEncoder(postData).Encode(txnReq)

	req, err = http.NewRequest("POST", fmt.Sprintf("https://www.demo.fx-scale.com/scale/v1/transactions-service/post-transactions/%v", clientID), /*strings.NewReader(string(postData))*/postData)
	if err != nil {
		log.Fatal(err)
	}

	currentTime = time.Now()

	service = "transactions-service"
	subService = "post-transactions"
	clientID = "EZL"
	transactionID = uuid.New() //"e0d63ba997354229a8e458e6b04d1ced"
	scaleTS = fmt.Sprintf("%v", currentTime.UnixNano())
	apiKey = "a1fc360d-fe95-4182-a9f3-a7180887771a"
	scaleHMAC = genHMAC(service, subService, clientID, scaleTS, transactionID, apiKey)

	req.Header.Add("SCALE_TS", scaleTS)
	req.Header.Add("SCALE_TX_ID", transactionID)
	req.Header.Add("SCALE_HMAC", scaleHMAC)
	req.Header.Add("Content-Type", "application/json")

	//fmt.Printf("TS: %v\n", req.Header.Get("SCALE_TS"))
	//fmt.Printf("TX: %v\n", req.Header.Get("SCALE_TX"))
	//fmt.Printf("HMAC: %v\n", req.Header.Get("SCALE_HMAC"))

	//if *quoteCmd == true {
	//	cmdline := "curl -k --header \"SCALE_TS: %v\" --header \"SCALE_TX_ID: %v\" --header \"SCALE_HMAC: %v\" --header \"Content-Type: application/json\" --cert \"%v:1234\" --tlsv1 \"https://www.demo.fx-scale.com/scale/v1/quotes-service/get-quotes/EZL/*\"\n"
	//	fmt.Printf(cmdline, scaleTS, transactionID, scaleHMAC, cnf.Scale.CertFile)
	//	return
	//}

	// Save a copy of this request for debugging.
	requestDump, err = httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Dump response
	/*
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))

		var txnResponse PostTransactionsResponse

		json.Unmarshal(data, &txnResponse)

		fmt.Printf("%v\n", txnResponse)
	*/

	var txnResponse PostTransactionsResponse
	json.NewDecoder(resp.Body).Decode(&txnResponse)
	fmt.Printf("%v\n", txnResponse)

}

func floatUnitsToMicros(floatUnits float64) int64 {
	return int64(floatUnits * 1000000.0)
}

func genHMAC(service string, subService string, clientID string, utcTime string, clientTxID string, apiKey string) string {
	message := fmt.Sprintf("%v:%v:%v:%v:%v", service, subService, clientID, utcTime, clientTxID)
	return computeHmac256(message, apiKey)
}

func computeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
