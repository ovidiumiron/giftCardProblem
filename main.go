package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ovidiumiron/giftCardProblem/findPair"
	"github.com/ovidiumiron/giftCardProblem/priceProduct"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatalf("Not enought number arguments in command line:%v\n", os.Args)
	}

	pricesProducts, err := priceProduct.ReadFromFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	toSpend, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// Adapt pricesProducts to the requested input for findPair.FindPair.
	data := make([]findPair.Interface, len(pricesProducts))
	for idx, d := range pricesProducts {
		data[idx] = d
	}

	s, err := findPair.FindPair(data, toSpend)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s, %s\n", s.First, s.Second)
}
