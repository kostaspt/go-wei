package wei

import (
	"math/big"

	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
)

type Wei big.Int

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
