package findPair

import (
	"fmt"
)

type InvalidSizeError []Interface

func (p InvalidSizeError) Error() string {
	return fmt.Sprintf("The size of the slice is <%d>. Minim accepted value is 2.", len(p))
}

type TooLessMoneyError int

func (t TooLessMoneyError) Error() string {
	return fmt.Sprintf("Not enoght money to buy any pressent. The best approximation is: %d", t)
}

type Interface interface {
	Price() int
	Product() string
}

type Pair struct {
	First  Interface
	Second Interface
}

func FindPair(data []Interface, realMoney int) (pair Pair, err error) {
	if l := len(data); l < 2 {
		err = InvalidSizeError(data)
		return
	}
	left, right := 0, len(data)-1
	needToSpend := data[left].Price() + data[right].Price()

	for left+1 < right {
		if needToSpend == realMoney {
			break
		}
		if needToSpend < realMoney {
			left += 1
		}
		if needToSpend > realMoney {
			right -= 1
		}
		needToSpend = data[left].Price() + data[right].Price()
	}

	if needToSpend > realMoney {
		err = TooLessMoneyError(needToSpend)
		return
	}

	return Pair{data[left], data[right]}, nil
}
