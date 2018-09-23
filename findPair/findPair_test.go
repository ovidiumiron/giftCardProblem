package findPair

import (
	"reflect"
	"testing"
)

type Concrete struct {
	foo int
	bar string
}

func (p Concrete) Price() int {
	return p.foo
}

func (p Concrete) Product() string {
	return p.bar
}
func Adapt(p []Concrete) []Interface {
	d := make([]Interface, len(p))
	for idx, c := range p {
		d[idx] = c
	}
	return d
}

func TestFindPair(t *testing.T) {
	var tests = []struct {
		slice []Concrete
		sum   int
		pair  Pair
		err   error
	}{
		{[]Concrete{}, 99, Pair{}, InvalidSizeError(Adapt([]Concrete{}))},

		{[]Concrete{Concrete{1, "product"}}, 99, Pair{}, InvalidSizeError(Adapt([]Concrete{Concrete{1, "product"}}))},
		// positive test
		{
			[]Concrete{Concrete{1, "product1"}, Concrete{2, "product2"}},
			3,
			Pair{Concrete{1, "product1"}, Concrete{2, "product2"}},
			nil,
		},
		// positive test
		{
			[]Concrete{Concrete{1, "product1"}, Concrete{2, "product2"}, Concrete{3, "product3"}},
			3,
			Pair{Concrete{1, "product1"}, Concrete{2, "product2"}},
			nil,
		},
		// basic positive test
		{
			[]Concrete{Concrete{1, "product1"}, Concrete{2, "product2"}, Concrete{3, "product3"}},
			3,
			Pair{Concrete{1, "product1"}, Concrete{2, "product2"}},
			nil,
		},
		// basic negativ test
		{
			[]Concrete{Concrete{10, "product1"}, Concrete{20, "product2"}, Concrete{30, "product3"}},
			3,
			Pair{},
			TooLessMoneyError(30),
		},
	}
	for _, test := range tests {
		got, err := FindPair(Adapt(test.slice), test.sum)
		if got != test.pair {
			t.Errorf("FindPair(%v, %v) = %+v expected %+v.", test.slice, test.sum, got, test.pair)
		}

		if reflect.DeepEqual(err, test.err) == false {
			t.Errorf("Solution(%v, %v) return error <%s> but expected <%s>.", test.slice, test.sum, err, test.err)
		}
	}
}
