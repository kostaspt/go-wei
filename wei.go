package wei

import (
	"bytes"
	"math/big"

	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
)

type Wei big.Int

func New(val uint64) Wei {
	b := big.Int{}
	return NewFromBigInt(b.SetUint64(val))
}

func NewFromBigInt(val *big.Int) Wei {
	return Wei(*val)
}

func (w Wei) Ether() decimal.Decimal {
	val := big.Int(w)
	return decimal.NewFromBigInt(&val, 1).Div(decimal.New(10, 18))
}

func (w Wei) BigInt() *big.Int {
	val := big.Int(w)
	return &val
}

func (w *Wei) Scan(src interface{}) error {
	parsed := new(pgtype.Numeric)
	if err := parsed.Scan(src); err != nil {
		return err
	}

	*w = Wei(*decimal.NewFromBigInt(parsed.Int, parsed.Exp).BigInt())

	return nil
}

func (w Wei) MarshalJSON() ([]byte, error) {
	val := big.Int(w)
	return val.MarshalJSON()
}

func (w *Wei) UnmarshalJSON(b []byte) error {
	b = bytes.ReplaceAll(b, []byte("\""), []byte{})

	bi := new(big.Int)

	if err := bi.UnmarshalJSON(b); err != nil {
		return err
	}

	*w = NewFromBigInt(bi)

	return nil
}
