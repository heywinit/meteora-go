package meteora

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// DLMM represents the core SDK structure.
type DLMM struct {
	Client   *rpc.Client
	Program  solana.PublicKey
	Network  string
}
