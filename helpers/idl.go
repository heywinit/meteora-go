package helpers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// FetchIDL fetches the program's IDL JSON from the Solana blockchain.
func FetchIDL(client *rpc.Client, programID solana.PublicKey) (map[string]interface{}, error) {
	data, err := client.GetAccountInfo(context.Background(), programID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch IDL: %w", err)
	}

	if data == nil || len(data.GetBinary()) == 0 {
		return nil, fmt.Errorf("IDL not found for program %s", programID.String())
	}

	var idlJSON map[string]interface{}
	err = json.Unmarshal(data.Value.Data.GetBinary(), &idlJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to parse IDL: %w", err)
	}

	return idlJSON, nil
}
