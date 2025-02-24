package meteora

import (
	"log"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/heywinit/meteora-go/helpers"
)

// DLMM represents the core SDK structure.
type DLMM struct {
	Client  *rpc.Client
	Program solana.PublicKey
	IDL     map[string]interface{}
}

// NewDLMM initializes DLMM with IDL.
func NewDLMM(network string) *DLMM {
	client := rpc.New("https://api.mainnet-beta.solana.com")
	program := solana.MustPublicKeyFromBase58("LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo")

	idl, err := helpers.FetchIDL(client, program)
	if err != nil {
		log.Println("Warning: Could not fetch IDL, using fallback")
		idl = map[string]interface{}{}
	}

	return &DLMM{
		Client:  client,
		Program: program,
		IDL:     idl,
	}
}