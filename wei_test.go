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

type testJSONStruct struct {
	Foo string `json:"foo"`
	Wei Wei    `json:"wei"`
}

func TestWei_MarshalJSON(t *testing.T) {
	val := testValue()

	m, err := json.Marshal(testJSONStruct{
		Foo: "bar",
		Wei: val,
	})
	assert.NoError(t, err)

	assert.Equal(t, `{"foo":"bar","wei":13370000000000000000}`, string(m))
}

func TestWei_UnmarshalJSON(t *testing.T) {
	var (
		res testJSONStruct
		j   []byte
		err error
	)

	j, err = json.Marshal(testJSONStruct{
		Foo: "bar",
		Wei: testValue(),
	})
	assert.NoError(t, err)

	err = json.Unmarshal(j, &res)
	assert.NoError(t, err)

	assert.Equal(t, "13.37", res.Wei.Ether().String())
	assert.Equal(t, "bar", res.Foo)
}

func testValue() Wei {
	val := big.NewInt(1337)

	// We want to get 13.37, so we need 2 fewer decimals than 18
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil)

	// Res: 13370000000000000000
	return NewFromBigInt(new(big.Int).Mul(val, exp))
}
