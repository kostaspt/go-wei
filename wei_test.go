package wei

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var val uint64 = 1337

	w := New(val)
	bi := new(big.Int).SetInt64(int64(val))

	assert.Equal(t, 0, bi.Cmp(w.BigInt()))
}

func TestWei_SetDecimals(t *testing.T) {
	val := testValue()

	val.SetDecimals(17)

	assert.True(t, val.Ether().Equal(decimal.NewFromFloat(133.7)))
}

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
	assert.NoError(t, err)

	assert.Equal(t, `{"foo":"bar","val":13370000000000000000}`, string(m))
}

func testValue() Wei {
	val := big.NewInt(1337)

	// We want to get 13.37, so we need 2 fewer decimals than 18
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil)

	// Res: 13370000000000000000
	return NewFromBigInt(new(big.Int).Mul(val, exp))
}
