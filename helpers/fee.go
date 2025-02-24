package helpers

import "math/big"

const (
	FeePrecision = 1_000_000_000
	MaxFeeRate   = 100_000_000
)

// ComputeFee calculates the transaction fee.
func ComputeFee(amount *big.Int, feeRate uint64) *big.Int {
	denominator := big.NewInt(int64(FeePrecision - feeRate))
	fee := new(big.Int).Mul(amount, big.NewInt(int64(feeRate)))
	return fee.Div(fee, denominator)
}
