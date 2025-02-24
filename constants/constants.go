package constants

import "github.com/gagliardetto/solana-go"

var (
	NetworkMainnet = "mainnet-beta"
	NetworkDevnet  = "devnet"
	NetworkLocal   = "localhost"

	LBCLMMProgramID = solana.MustPublicKeyFromBase58("LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo")
	AdminDevnet     = solana.MustPublicKeyFromBase58("6WaLrrRfReGKBYUSkmx2K6AuT21ida4j8at2SUiZdXu8")

	BasisPointMax = uint64(10000)
	ScaleOffset   = uint64(64)
)
