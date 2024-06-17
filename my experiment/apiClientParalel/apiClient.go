package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	urls := []string{"https://api.sampleapis.com/futurama/info",
		"https://www.cbr-xml-daily.ru/latest.js",
		"https://api.sampleapis.com/coffee/hot",
		"https://api.coindesk.com/v1/bpi/currentprice.json",
		"https://www.boredapi.com/api/activity",
		"https://api.nationalize.io/?name=nathaniel",
		"https://datausa.io/api/data?drilldowns=Nation&measures=Population",
		"https://randomuser.me/api/",
		"https://api.zippopotam.us/us/33162",
	}

	for index, url := range urls {
		wg.Add(1)
		go func(i int, apiUrl string) {
			defer wg.Done()
			getInfoFromApi(i, apiUrl)
		}(index, url)
	}
	wg.Wait()
}

func getInfoFromApi(i int, url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", url, err)
		return
	}
	fmt.Printf("%d Status: %s url= %s \n", i, res.Status, url)
	defer res.Body.Close()
}
