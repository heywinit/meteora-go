# Meteora Go SDK

A GoLang port of the Meteora SDK for building applications on top of **Dynamic CLMM (DLMM)**.

## 📌 Features
- Interact with **Meteora's Dynamic CLMM (DLMM)**
- Manage **liquidity pools**, **swaps**, and **fees**
- Connect to the **Solana blockchain**
- Lightweight and optimized for Go applications

## 🚀 Installation
```sh
go get github.com/byedeep/meteora-go
```

## Project Structure
```
meteora-go/
│── clmm/           # Dynamic CLMM logic
│── client/         # Solana client
│── utils/          # Helpers
│── types/          # Core types
│── internal/       # Internal utilities (not exported)
│── examples/       # Usage examples
│── meteora.go      # Main entry file
│── go.mod
│── README.md
```

## Quick Start
```
package main

import (
    "fmt"
    "github.com/byedeep/meteora-go-sdk/pkg/dlmm"
)
 
func main() {
    pool := clmm.NewPool()
    fmt.Println("Meteora Go SDK initialized:", pool)
}
```

## Running Tests
```
go test ./tests/...
```

## License:
MIT