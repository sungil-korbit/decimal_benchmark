package benches

import (
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"testing"
)

type OP string

const (
	ADD = OP("add")
	SUB = OP("sub")
	MUL = OP("mul")
	DIV = OP("div")
)

func decimalOp(op OP, x decimal.Decimal, y decimal.Decimal) {
	switch op {
	case ADD:
		x.Add(y)
	case SUB:
		x.Sub(y)
	case MUL:
		x.Mul(y)
	case DIV:
		x.Div(y)
	default:
	}
}

func foldOp(op OP, init decimal.Decimal, count int) {
	acc := init

	values := lo.RepeatBy[decimal.Decimal](count, func(i int) decimal.Decimal {
		return decimal.NewFromInt(int64(i + 1))
	})

	switch op {
	case ADD:
		for _, v := range values {
			acc.Add(v)
		}
	case SUB:
		for _, v := range values {
			acc.Sub(v)
		}
	case MUL:
		for _, v := range values {
			acc.Mul(v)
		}
	case DIV:
		for _, v := range values {
			acc.Div(v)
		}
	}
}

func BenchmarkAdd10k(b *testing.B) {
	foldOp(ADD, decimal.RequireFromString("0"), 10000)
}
func BenchmarkAddSelf(b *testing.B) {
	decimalOp(ADD, decimal.RequireFromString("2.01"), decimal.RequireFromString("2.01"))
}
func BenchmarkAddSimple(b *testing.B) {
	decimalOp(ADD, decimal.RequireFromString("2"), decimal.RequireFromString("2"))
}
func BenchmarkAddOne(b *testing.B) {
	decimalOp(ADD, decimal.RequireFromString("2.01"), decimal.RequireFromString("1"))
}
func BenchmarkAddTwo(b *testing.B) {
	decimalOp(ADD, decimal.RequireFromString("2.01"), decimal.RequireFromString("2"))
}
func BenchmarkAddOneHundred(b *testing.B) {
	decimalOp(ADD, decimal.RequireFromString("2.01"), decimal.RequireFromString("100"))
}
func BenchmarkAddPointZeroOne(b *testing.B) {
	decimalOp(ADD, decimal.RequireFromString("2.01"), decimal.RequireFromString("0.01"))
}
func BenchmarkAddNegativePointFive(b *testing.B) {
	decimalOp(ADD, decimal.RequireFromString("2.01"), decimal.RequireFromString("-0.5"))
}
func BenchmarkAddPi(b *testing.B) {
	decimalOp(ADD, decimal.RequireFromString("2.01"), decimal.RequireFromString("3.1415926535897932384626433832"))
}
func BenchmarkAddNegativePi(b *testing.B) {
	decimalOp(ADD, decimal.RequireFromString("2.01"), decimal.RequireFromString("-3.1415926535897932384626433832"))
}

func BenchmarkSub10k(b *testing.B) {
	foldOp(SUB, decimal.RequireFromString("5000000"), 10000)
}
func BenchmarkSubSelf(b *testing.B) {
	decimalOp(SUB, decimal.RequireFromString("2.01"), decimal.RequireFromString("2.01"))
}
func BenchmarkSubSimple(b *testing.B) {
	decimalOp(SUB, decimal.RequireFromString("2"), decimal.RequireFromString("2"))
}
func BenchmarkSubOne(b *testing.B) {
	decimalOp(SUB, decimal.RequireFromString("2.01"), decimal.RequireFromString("1"))
}
func BenchmarkSubTwo(b *testing.B) {
	decimalOp(SUB, decimal.RequireFromString("2.01"), decimal.RequireFromString("2"))
}
func BenchmarkSubOneHundred(b *testing.B) {
	decimalOp(SUB, decimal.RequireFromString("2.01"), decimal.RequireFromString("100"))
}
func BenchmarkSubPointZeroOne(b *testing.B) {
	decimalOp(SUB, decimal.RequireFromString("2.01"), decimal.RequireFromString("0.01"))
}
func BenchmarkSubNegativePointFive(b *testing.B) {
	decimalOp(SUB, decimal.RequireFromString("2.01"), decimal.RequireFromString("-0.5"))
}
func BenchmarkSubPi(b *testing.B) {
	decimalOp(SUB, decimal.RequireFromString("2.01"), decimal.RequireFromString("3.1415926535897932384626433832"))
}
func BenchmarkSubNegativePi(b *testing.B) {
	decimalOp(SUB, decimal.RequireFromString("2.01"), decimal.RequireFromString("-3.1415926535897932384626433832"))
}

/* DIV */
func BenchmarkDiv10k(b *testing.B) {
	foldOp(DIV, decimal.RequireFromString("1000000000000"), 10000)
}
func BenchmarkDivSelf(b *testing.B) {
	decimalOp(DIV, decimal.RequireFromString("2.01"), decimal.RequireFromString("2.01"))
}
func BenchmarkDivSimple(b *testing.B) {
	decimalOp(DIV, decimal.RequireFromString("2"), decimal.RequireFromString("2"))
}
func BenchmarkDivOne(b *testing.B) {
	decimalOp(DIV, decimal.RequireFromString("2.01"), decimal.RequireFromString("1"))
}
func BenchmarkDivTwo(b *testing.B) {
	decimalOp(DIV, decimal.RequireFromString("2.01"), decimal.RequireFromString("2"))
}
func BenchmarkDivOneHundred(b *testing.B) {
	decimalOp(DIV, decimal.RequireFromString("2.01"), decimal.RequireFromString("100"))
}
func BenchmarkDivPointZeroOne(b *testing.B) {
	decimalOp(DIV, decimal.RequireFromString("2.01"), decimal.RequireFromString("0.01"))
}
func BenchmarkDivNegativePointFive(b *testing.B) {
	decimalOp(DIV, decimal.RequireFromString("2.01"), decimal.RequireFromString("-0.5"))
}
func BenchmarkDivPi(b *testing.B) {
	decimalOp(DIV, decimal.RequireFromString("2.01"), decimal.RequireFromString("3.1415926535897932384626433832"))
}
func BenchmarkDivNegativePi(b *testing.B) {
	decimalOp(DIV, decimal.RequireFromString("2.01"), decimal.RequireFromString("-3.1415926535897932384626433832"))
}
func BenchmarkDivUnderflow(b *testing.B) {
	decimalOp(DIV, decimal.RequireFromString("1.02343545345"), decimal.RequireFromString("0.35454343453"))
}

/* MUL */
func BenchmarkMul10k(b *testing.B) {
	foldOp(MUL, decimal.RequireFromString("1.1"), 25)
}
func BenchmarkMulSelf(b *testing.B) {
	decimalOp(MUL, decimal.RequireFromString("2.01"), decimal.RequireFromString("2.01"))
}
func BenchmarkMulSimple(b *testing.B) {
	decimalOp(MUL, decimal.RequireFromString("2"), decimal.RequireFromString("2"))
}
func BenchmarkMulOne(b *testing.B) {
	decimalOp(MUL, decimal.RequireFromString("2.01"), decimal.RequireFromString("1"))
}
func BenchmarkMulTwo(b *testing.B) {
	decimalOp(MUL, decimal.RequireFromString("2.01"), decimal.RequireFromString("2"))
}
func BenchmarkMulOneHundred(b *testing.B) {
	decimalOp(MUL, decimal.RequireFromString("2.01"), decimal.RequireFromString("100"))
}
func BenchmarkMulPointZeroOne(b *testing.B) {
	decimalOp(MUL, decimal.RequireFromString("2.01"), decimal.RequireFromString("0.01"))
}
func BenchmarkMulNegativePointFive(b *testing.B) {
	decimalOp(MUL, decimal.RequireFromString("2.01"), decimal.RequireFromString("-0.5"))
}
func BenchmarkMulPi(b *testing.B) {
	decimalOp(MUL, decimal.RequireFromString("2.01"), decimal.RequireFromString("3.1415926535897932384626433832"))
}
func BenchmarkMulNegativePi(b *testing.B) {
	decimalOp(MUL, decimal.RequireFromString("2.01"), decimal.RequireFromString("-3.1415926535897932384626433832"))
}
