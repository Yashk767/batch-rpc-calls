package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	client, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	//ethClient := ethclient.NewClient(client)
	ctx := context.Background()

	// The contract ABI as a JSON string (simplified for example)
	contractABIJSON := `[{"constant":true,"inputs":[],"name":"getNumber","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},
                         {"constant":true,"inputs":[],"name":"getName","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"}]`

	contractABI, err := abi.JSON(strings.NewReader(contractABIJSON))
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")

	// Pack the data for the 'getNumber' function
	getData, err := contractABI.Pack("getNumber")
	if err != nil {
		log.Fatal(err)
	}

	// Pack the data for the 'getName' function
	getNameData, err := contractABI.Pack("getName")
	if err != nil {
		log.Fatal(err)
	}

	var resultNumber string // use string to capture the hex result
	var resultString string

	// Batch call setup
	call1 := rpc.BatchElem{
		Method: "eth_call",
		Args:   []interface{}{map[string]interface{}{"to": contractAddress.Hex(), "data": fmt.Sprintf("%x", getData)}, "latest"},
		Result: &resultNumber,
	}
	call2 := rpc.BatchElem{
		Method: "eth_call",
		Args:   []interface{}{map[string]interface{}{"to": contractAddress.Hex(), "data": fmt.Sprintf("%x", getNameData)}, "latest"},
		Result: &resultString,
	}

	// Perform batch RPC call
	err = client.BatchCallContext(ctx, []rpc.BatchElem{call1, call2})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result from getNumber in hex:", *(call1.Result.(*string)))
	fmt.Println("Result from getName in hex:", *(call2.Result.(*string)))

	if resultNumber != "" {
		// Process the number result
		num := new(big.Int)
		num.SetString(strings.TrimPrefix(resultNumber, "0x"), 16)
		fmt.Println("Decoded result from getNumber:", num)
	}

	// Assuming 'resultString' is your hex-encoded string from an RPC call
	resultName := *call2.Result.(*string) // This should be set from your RPC call results

	// Remove the "0x" prefix and decode
	cleanHex := strings.TrimPrefix(resultName, "0x")
	decodedBytes, err := hex.DecodeString(cleanHex)
	if err != nil {
		log.Fatal("Failed to decode hex string:", err)
	}

	// Convert byte slice to string
	decodedString := string(decodedBytes)
	fmt.Println("Decoded result from getName:", decodedString)

}
