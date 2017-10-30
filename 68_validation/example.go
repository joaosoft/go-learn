package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type Test struct {
	A decimal.Decimal `json:"a"`
	B decimal.Decimal `json:"b"`
}

// IsValid validates the object's invariants.
func (test *Test) IsValid() bool {
	return test.hasValidA() &&
		test.hasValidB() &&
		test.B.Cmp(test.A) != 1
}

func (test *Test) hasValidA() bool {
	return test.A.Cmp(decimal.Zero) == 1
}

func (test *Test) hasValidB() bool {
	return test.B.Cmp(decimal.Zero) == 1
}

func NewTest(a decimal.Decimal, b decimal.Decimal) (*Test, error) {
	test := new(Test)
	test.A = a
	test.B = b

	if !test.IsValid() {
		return nil, fmt.Errorf("error!")
	}

	return test, nil
}

func main() {
	negativeDecimal, _ := decimal.NewFromString("-0.01")
	zeroDecimal, _ := decimal.NewFromString("0.00")
	positiveDecimal, _ := decimal.NewFromString("0.01")
	biggerDecimal, _ := decimal.NewFromString("0.02")

	var test *Test
	var err error

	test, err = NewTest(negativeDecimal, negativeDecimal)
	fmt.Println("NEGATIVE ( TEST:", test, " ; ERROR:", err)

	test, err = NewTest(zeroDecimal, zeroDecimal)
	fmt.Println("ZERO ( TEST:", test, " ; ERROR:", err)

	test, err = NewTest(positiveDecimal, positiveDecimal)
	fmt.Println("POSITIVE ( TEST:", test, " ; ERROR:", err)

	test, err = NewTest(biggerDecimal, biggerDecimal)
	fmt.Println("BIGGER ( TEST:", test, " ; ERROR:", err)

	test, err = NewTest(biggerDecimal, positiveDecimal)
	fmt.Println("A > B ( TEST:", test, " ; ERROR:", err)

	test, err = NewTest(positiveDecimal, biggerDecimal)
	fmt.Println("B > A ( TEST:", test, " ; ERROR:", err)

	test, err = NewTest(positiveDecimal, negativeDecimal)
	fmt.Println("POSITIVE | NEGATIVE ( TEST:", test, " ; ERROR:", err)

	test, err = NewTest(positiveDecimal, zeroDecimal)
	fmt.Println("POSITIVE | ZERO ( TEST:", test, " ; ERROR:", err)

}
