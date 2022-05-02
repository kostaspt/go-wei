package wei

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestWei_Ether(t *testing.T) {
	val := testValue()

	assert.True(t, val.Ether().Equal(decimal.NewFromFloat(13.37)))
}

func TestWei_MarshalJSON(t *testing.T) {
	val := testValue()

	m, err := json.Marshal(map[string]interface{}{
		"foo": "bar",
		"val": val,
	})
	if err != nil {
		return
	}

	assert.Equal(t, `{"foo":"bar","val":13370000000000000000}`, string(m))
}

func testValue() Wei {
	val := big.NewInt(1337)

	// We want to get 13.37, so we need 2 fewer decimals than 18
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil)

	// Res: 13370000000000000000
	return NewFromBigInt(new(big.Int).Mul(val, exp))
}
