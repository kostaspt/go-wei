package wei

import (
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestWei_Ether(t *testing.T) {
	val := big.NewInt(1337)

	// We want to get 13.37, so we need 2 fewer decimals than 18
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil)

	// Res: 13370000000000000000
	w := NewFromBigInt(new(big.Int).Mul(val, exp))

	assert.True(t, w.Ether().Equal(decimal.NewFromFloat(13.37)))
}
