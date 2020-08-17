package main

import (
	"fmt"
	"github.com/giantliao/beatles-protocol/price"
)

func main() {
	p, err := price.GetPrice()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*p)

}
