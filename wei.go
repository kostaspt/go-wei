package wei

import (
	"encoding/json"
	"math/big"

	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
)

type Wei struct {
	raw      *big.Int
	decimals uint8
}

type jsonWei struct {
	Value    string   `json:"value"`
	Raw      *big.Int `json:"raw"`
	Decimals uint8    `json:"decimals"`
}

func New(val uint64, decimals uint8) Wei {
	return Wei{
		raw:      new(big.Int).SetUint64(val),
		decimals: decimals,
	}
}

func NewFromBigInt(val *big.Int, decimals uint8) Wei {
	return Wei{
		raw:      val,
		decimals: decimals,
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
	j := jsonWei{
		Value:    w.Ether().String(),
		Raw:      w.raw,
		Decimals: w.decimals,
	}

	return json.Marshal(j)
}

func (w *Wei) UnmarshalJSON(b []byte) error {
	var (
		res jsonWei
		err error
	)

	if err = json.Unmarshal(b, &res); err != nil {
		return err
	}

	*w = Wei{
		raw:      res.Raw,
		decimals: res.Decimals,
	}

	return nil
}
