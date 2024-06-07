package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://galmart.kz/api/v1/catalog/shops")

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

}
