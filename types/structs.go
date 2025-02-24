package types

import (
	"math/big"

	"github.com/gagliardetto/solana-go"
)

type Decimal struct {
	Value string
}

// ClmmProgram represents the DLMM Solana program.
type ClmmProgram struct {
	ProgramID solana.PublicKey
}

// LbPair represents the on-chain liquidity bin pair.
type LbPair struct {
	TokenX solana.PublicKey
	TokenY solana.PublicKey
}

// LbPairAccount is a wrapper for the account.
type LbPairAccount struct {
	PublicKey solana.PublicKey
	Account   LbPair
}

// FeeInfo defines fee-related parameters.
type FeeInfo struct {
	BaseFeeRatePercentage string
	MaxFeeRatePercentage  string
	ProtocolFeePercentage string
}

// BinAndAmount represents bin allocations.
type BinAndAmount struct {
	BinID             int
	XAmountBpsOfTotal *big.Int
	YAmountBpsOfTotal *big.Int
}

// TokenReserve holds details about a token reserve.
type TokenReserve struct {
	PublicKey solana.PublicKey
	Reserve   solana.PublicKey
	Amount    *big.Int
	Decimal   int16
}

type LbPosition struct {
	PublicKey    solana.PublicKey
	PositionData PositionData
	Version      PositionVersion
}

type PositionInfo struct {
	PublicKey           solana.PublicKey
	LBPair              LbPair
	TokenX              TokenReserve
	TokenY              TokenReserve
	LBPairPositionsData []LbPosition
}

type PositionData struct {
	TotalXAmount           string
	TotalYAmount           string
	PositionBinData        []PositionBinData
	LastUpdatedAt          big.Int
	UpperBinId             int16
	LowerBinId             int16
	FeeX                   big.Int
	FeeY                   big.Int
	RewardOne              big.Int
	RewardTwo              big.Int
	FeeOwner               solana.PublicKey
	TotalClaimedFeeXAmount big.Int
	TotalClaimedFeeYAmount big.Int
}

type PositionBinData struct {
	BinId             int16
	Price             string
	PricePerToken     string
	BinXAmount        string
	BinYAmount        string
	BinLiquidity      string
	PositionLiquidity string
	PositionXAmount   string
	PositionYAmount   string
}

type EmissionRate struct {
	RewardOne Decimal
	RewardTwo Decimal
}

type SwapFee struct {
	FeeX big.Int
	FeeY big.Int
}

type LMRewards struct {
	RewardOne Decimal
	RewardTwo Decimal
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

type PairType int

const (
	Permissionless PairType = iota
	Permissioned
)

type ActivationType int

const (
	Slot ActivationType = iota
	Timestamp
)

type Strategy struct {
	SpotBalanced     map[string]interface{}
	CurveBalanced    map[string]interface{}
	BidAskBalanced   map[string]interface{}
	SpotImBalanced   map[string]interface{}
	CurveImBalanced  map[string]interface{}
	BidAskImBalanced map[string]interface{}
}

var strategy = Strategy{
	SpotBalanced:     map[string]interface{}{"spotBalanced": struct{}{}},
	CurveBalanced:    map[string]interface{}{"curveBalanced": struct{}{}},
	BidAskBalanced:   map[string]interface{}{"bidAskBalanced": struct{}{}},
	SpotImBalanced:   map[string]interface{}{"spotImBalanced": struct{}{}},
	CurveImBalanced:  map[string]interface{}{"curveImBalanced": struct{}{}},
	BidAskImBalanced: map[string]interface{}{"bidAskImBalanced": struct{}{}},
}

type StrategyParameters struct {
	MaxBinId     int
	MinBinId     int
	StrategyType StrategyType
	SingleSidedX *bool
}

type TQuoteCreatePositionParams struct {
	Strategy StrategyParameters
}

type TInitializePositionAndAddLiquidityParams struct {
	PositionPubKey       solana.PublicKey
	TotalXAmount         *big.Int
	TotalYAmount         *big.Int
	XYAmountDistribution []BinAndAmount
	User                 solana.PublicKey
	Slippage             *float64
}

type TInitializePositionAndAddLiquidityParamsByStrategy struct {
	PositionPubKey solana.PublicKey
	TotalXAmount   *big.Int
	TotalYAmount   *big.Int
	Strategy       StrategyParameters
	User           solana.PublicKey
	Slippage       *float64
}

type BinLiquidity struct {
	binId         int
	xAmount       big.Int
	yAmount       big.Int
	supply        big.Int
	version       int
	price         string
	pricePerToken string
}

/*

# export module BinLiquidity {

# export function empty(

*/

type SwapQuote struct {
	consumedInAmount big.Int
	outAmount        big.Int
	fee              big.Int
	protocolFee      big.Int
	minOutAmount     big.Int
	priceImpact      big.Int
	binArraysPubkey  []any
	endPrice         Decimal
}

type SwapQuoteExactOut struct {
	inAmount        big.Int
	outAmount       big.Int
	fee             big.Int
	priceImpact     Decimal
	protocolFee     big.Int
	maxInAmount     big.Int
	binArraysPubkey []any
}

// type AccountsCache struct {
// 	BinArrays map[string]BinArray
// 	LbPair    LbPair
// }

type SwapWithPriceImpactParams struct {
	/**
	 * mint of in token
	 */
	inToken solana.PublicKey
	/**
	 * mint of out token
	 */
	outToken solana.PublicKey
	/**
	 * in token amount
	 */
	inAmount big.Int
	/**
	 * price impact in bps
	 */
	priceImpact big.Int
	/**
	 * desired lbPair to swap against
	 */
	lbPair solana.PublicKey
	/**
	 * user
	 */
	user            solana.PublicKey
	binArraysPubkey []solana.PublicKey
}

type SwapParams struct {
	/**
	 * mint of in token
	 */
	inToken solana.PublicKey
	/**
	 * mint of out token
	 */
	outToken solana.PublicKey
	/**
	 * in token amount
	 */
	inAmount big.Int
	/**
	 * minimum out with slippage
	 */
	minOutAmount big.Int
	/**
	 * desired lbPair to swap against
	 */
	lbPair solana.PublicKey
	/**
	 * user
	 */
	user            solana.PublicKey
	binArraysPubkey []solana.PublicKey
}

type SwapExactOutParams struct {
	/**
	 * mint of in token
	 */
	inToken solana.PublicKey
	/**
	 * mint of out token
	 */
	outToken solana.PublicKey
	/**
	 * out token amount
	 */
	outAmount big.Int
	/**
	 * maximum in amount, also known as slippage
	 */
	maxInAmount big.Int
	/**
	 * desired lbPair to swap against
	 */
	lbPair solana.PublicKey
	/**
	 * user
	 */
	user            solana.PublicKey
	binArraysPubkey []solana.PublicKey
}

// type GetOrCreateATAResponse struct{

type BitmapType int

const (
	U1024 BitmapType = iota
	U512
)
