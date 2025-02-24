package helpers

import (
	"context"
	"log"

	"github.com/gagliardetto/solana-go/rpc"
)

var SolanaClient *rpc.Client

func InitRPCClient(endpoint string) {
	SolanaClient = rpc.New(endpoint)
	log.Println("Connected to Solana RPC:", endpoint)
}

// Example: Fetch latest block height
func GetBlockHeight() uint64 {
	height, err := SolanaClient.GetSlot(
		context.Background(),
		rpc.CommitmentFinalized,
	)
	if err != nil {
		log.Fatalf("Error fetching block height: %v", err)
	}
	return height
}
