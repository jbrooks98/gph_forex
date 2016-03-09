package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
        "os"
	"log"
)

type Currency struct {
	Terms string
	Privacy string
	Timestamp int
	Source string
	Quotes map[string]float64
 }

func main() {

	// url will return
	//{
        //"terms": "https://currencylayer.com/terms",
        //"privacy": "https://currencylayer.com/privacy",
        //"timestamp": 1430401802,
        //"source": "USD",
        //"quotes": {
        //"USDAED": 3.672982,
        //"USDAFN": 57.8936
        //}
        //}

	currency_url := "http://apilayer.net/api/"
	live_endpoint := "live"
	access_key := "?access_key="
        api_key := os.Getenv("CURRENCY_API_KEY")
        live_response := currency_url + live_endpoint + access_key + api_key

	resp, err := http.Get(live_response)

	defer resp.Body.Close()

	if err != nil {
                 panic(err)
         }

	// read json http response
	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
                 log.Fatal("http response broke: %s\n", err)
	}

	currency_response := Currency{}

	err = json.Unmarshal([]byte(jsonDataFromHttp), &currency_response)

	if err != nil {
		log.Fatalf("shit broke: %s\n", err)
	}

	// test struct data
	fmt.Println(currency_response)
}
