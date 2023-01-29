package big_sense

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/big"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// ValueOf takes int or float  and returns a new Int sets to integer part of the number.
func ValueOf[T Number](number T) *big.Int {
	return Empty().SetInt64(int64(number))
}

// Val is allias of ValueOf
func Val[T Number](number T) *big.Int {
	return ValueOf(number)
}

func FromStringBase(number string, base int) (*big.Int, error) {
	if base < 2 {
		return nil, fmt.Errorf("big_sense: base can-not be less than 2. base %v", base)
	}
	n := new(big.Int)
	if n, ok := n.SetString(number, base); ok {
		return n, nil
	}
	return nil, fmt.Errorf("big_sense: '%v' convertion can-not be done", number)
}

func FromString(number string) (*big.Int, error) {
	return FromStringBase(number, 10)
}

// Atobi ASCII to Big Int. Allias of FromString.
func Atobi(number string) (*big.Int, error) {
	return FromString(number)
}

func Empty() *big.Int {
	return new(big.Int)
}

// Zero returns immediate 0 value big.Int
func Zero() *big.Int {
	return ValueOf(0)
}

// One returns immediate 1 value big.Int
func One() *big.Int {
	return ValueOf(1)
}

func Add(first *big.Int, second *big.Int) *big.Int {
	return Empty().Add(first, second)
}

func AddAll(first *big.Int, rest ...*big.Int) *big.Int {
	sum := first
	for _, rest := range rest {
		sum = sum.Add(sum, rest)
	}
	return sum
}

func Sub(first *big.Int, second *big.Int) *big.Int {
	return Empty().Sub(first, second)
}

func Mul(first *big.Int, second *big.Int) *big.Int {
	return Empty().Mul(first, second)
}

func MulAll(first *big.Int, rest ...*big.Int) *big.Int {
	total := first
	for _, rest := range rest {
		total = Mul(total, rest)
	}
	return total
}

func MulRange(a, b int64) *big.Int {
	return Empty().MulRange(a, b)
}

func Div(dividend *big.Int, divisor *big.Int) *big.Int {
	return Empty().Div(dividend, divisor)
}

func Sqrt(number *big.Int) *big.Int {
	return Empty().Sqrt(number)
}

func Cmp(x, y *big.Int) int {
	return x.Cmp(y)
}
