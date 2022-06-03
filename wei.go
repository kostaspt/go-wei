package wei

import (
	"bytes"
	"math/big"

	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
)

type Wei struct {
	raw      *big.Int
	decimals uint8
}

func New(val uint64) Wei {
	b := big.Int{}
	return NewFromBigInt(b.SetUint64(val))
}

func NewFromBigInt(val *big.Int) Wei {
	return Wei{
		raw:      val,
		decimals: 18,
	}
}

func (w *Wei) SetDecimals(d uint8) {
	w.decimals = d
}

func (w Wei) Ether() decimal.Decimal {
	return decimal.
		NewFromBigInt(w.BigInt(), 1).
		Div(decimal.New(10, int32(w.decimals)))
}

func (w Wei) BigInt() *big.Int {
	return w.raw
}

func (w *Wei) Scan(src interface{}) error {
	parsed := new(pgtype.Numeric)
	if err := parsed.Scan(src); err != nil {
		return err
	}

	*w = Wei{raw: decimal.NewFromBigInt(parsed.Int, parsed.Exp).BigInt()}

	return nil
}

func (w Wei) MarshalJSON() ([]byte, error) {
	return w.raw.MarshalJSON()
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
