package main

import (
	"fmt"
	"struct/internal"
)

func main() {
	cust := internal.Customer{
		Age:      44,
		Balance:  10000,
		Debt:     1000,
		Name:     "vetrof",
		Discount: true,
	}

	fmt.Println(cust.WrOffDebt())
	fmt.Println(cust.CalcDiscount())
	fmt.Printf("%+v\n", cust)

}
