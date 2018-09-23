package priceProduct

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LineBadFormatError string

func (l LineBadFormatError) Error() string {
	return fmt.Sprintf("The line %q hasn't have the format <present, value>.", l)
}

type PriceNotIntError string

func (p PriceNotIntError) Error() string {
	return fmt.Sprintf("The price %q can not be converted to int.", p)
}

func ReadFromFile(path string) ([]PriceProduct, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	pricesProducts := []PriceProduct{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		if len(s) != 2 {
			return nil, LineBadFormatError(line)
		}

		product, price := strings.TrimSpace(s[0]), strings.TrimSpace(s[1])
		intPrice, err := strconv.Atoi(price)
		if err != nil {
			return nil, PriceNotIntError(price)
		}

		priceProduct, err := NewPriceProduct(intPrice, product)
		if err != nil {
			return nil, err
		}

		pricesProducts = append(pricesProducts, priceProduct)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return pricesProducts, nil
}
