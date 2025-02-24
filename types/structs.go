package types

import (
	"math/big"

	"github.com/gagliardetto/solana-go"
)

// FeeInfo defines fee-related parameters.
type FeeInfo struct {
	BaseFeeRatePercentage string
	MaxFeeRatePercentage  string
	ProtocolFeePercentage string
}

// BinAndAmount represents bin allocations.
type BinAndAmount struct {
	BinID              int
	XAmountBpsOfTotal  *big.Int
	YAmountBpsOfTotal  *big.Int
}

// TokenReserve holds details about a token reserve.
type TokenReserve struct {
	PublicKey solana.PublicKey
	Reserve   solana.PublicKey
	Amount    *big.Int
	Decimal  int16
}

type LbPosition struct {
	PublicKey solana.PublicKey
	PositionData PositionData
	Version PositionVersion
}

type PositionData struct {
	TotalXAmount string
	TotalYAmount string
	PositionBinData []PositionBinData
	LastUpdatedAt big.Int
	UpperBinId int16
	LowerBinId int16
	FeeX big.Int
	FeeY big.Int
	RewardOne big.Int
	RewardTwo big.Int
	FeeOwner solana.PublicKey
	TotalClaimedFeeXAmount big.Int
	TotalClaimedFeeYAmount big.Int
}

type PositionBinData struct {
	BinId int16
	Price string
	PricePerToken string
	BinXAmount string;
	BinYAmount string;
	BinLiquidity string;
	PositionLiquidity string;
	PositionXAmount string;
	PositionYAmount string;
}

// StrategyType Enum
type StrategyType int

const (
	SpotImBalanced StrategyType = iota
	CurveImBalanced
	BidAskImBalanced
	SpotBalanced
	CurveBalanced
	BidAskBalanced
)

// PositionVersion Enum
type PositionVersion int

const (
	V1 PositionVersion = iota
	V2
)
