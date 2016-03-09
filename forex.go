package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
        "os"
)

type Currency struct {
	terms string
	privacy string
	timestamp int
	source string
	quotes map[string]float64
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
        //"USDAFN": 57.8936,
        //}
        //}

	currency_url := "http://apilayer.net/api/"
	var live_endpoint string = "live"
	var access_key string = "?access_key=
        var api_key string = os.Getenv("CURRENCY_API_KEY")
        var live_response string = currency_url + live_endpoint + access_key + api_key

	resp, err := http.Get(live_response)

	defer resp.Body.Close()

	if err != nil {
                 panic(err)
         }

	// read json http response
	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
                 panic(err)
	}

	var jsonData []Currency

	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)

	if err != nil {
		panic(err)
	}

	// test struct data
	fmt.Println(jsonData)
}
