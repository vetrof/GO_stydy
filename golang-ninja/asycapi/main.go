package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	go apiService("https://api.sampleapis.com/beers/ale")
	go apiService("https://api.sampleapis.com/beers/stouts")
	go apiService("https://api.sampleapis.com/beers/red-ale")
	go apiService("https://api.sampleapis.com/coffee/hot")
	go apiService("https://api.sampleapis.com/coffee/iced")
	go apiService("https://api.sampleapis.com/futurama/info")
	go apiService("https://api.sampleapis.com/futurama/characters")
	go apiService("https://api.sampleapis.com/futurama/cast")
	go apiService("https://api.sampleapis.com/futurama/episodes")
	go apiService("https://api.sampleapis.com/futurama/questions")
	go apiService("https://api.sampleapis.com/futurama/inventory")
	go apiService("https://www.cbr-xml-daily.ru/latest.js")

	time.Sleep(5 * time.Second)

}

func apiService(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer response.Body.Close()

	//body, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	fmt.Println(response.Status, url)
}
