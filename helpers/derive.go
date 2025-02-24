package helpers

import (
	"github.com/gagliardetto/solana-go"
)

// DeriveLbPair derives the PDA for a Liquidity Bin Pair (LB Pair)
func DeriveLbPair(tokenX, tokenY solana.PublicKey, programID solana.PublicKey) (solana.PublicKey, error) {
	keys := [][]byte{tokenX.Bytes(), tokenY.Bytes()}
	pda, _, err := solana.FindProgramAddress(keys, programID)
	return pda, err
}
