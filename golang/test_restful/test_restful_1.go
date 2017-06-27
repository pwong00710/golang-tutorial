package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v0"
)

type country struct {
	ID          int    `json:"id"`
	CountryName string `json:"countryName"`
	Population  int    `json:"population"`
}

func main() {
	// GET request
	resp, err := resty.R().Get("http://192.168.0.101:8080/rest/getAllCountries")

	// explore response object
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v [%T]", resp, resp) // or resp.String() or string(resp.Body())
	//fmt.Printf("\nResponse Body: %v [%T]", resp.Result(), resp.Result())
	/*
		var countryMap []map[string]interface{}
		err = json.Unmarshal(resp.Body(), &countryMap)

		if err != nil {
			fmt.Printf("\nError: %v", err)
		} else {
			fmt.Printf("\nSize: %v", len(countryMap))
		}

		var countryList []country

		for _, data := range countryMap {
			var c country
			c.id, _ = strconv.Atoi(fmt.Sprintf("%v", data["id"]))
			c.countryName = fmt.Sprintf("%s", data["countryName"])
			c.population, _ = strconv.Atoi(fmt.Sprintf("%v", data["population"]))
			countryList = append(countryList, c)
		}

		fmt.Println(countryList)
	*/
	var countryList []country
	err = json.Unmarshal(resp.Body(), &countryList)

	if err != nil {
		fmt.Printf("\nError: %v", err)
	}

	fmt.Println(countryList)
}
