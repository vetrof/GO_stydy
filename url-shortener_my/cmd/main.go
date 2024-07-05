package main

import (
	"fmt"
	"url-shortener_my/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
