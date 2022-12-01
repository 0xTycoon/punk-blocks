package main

import (
	"encoding/hex"
	"fmt"
	"github.com/miguelmota/go-solidity-sha3"
	"testing"
)

func TestThis(t *testing.T) {
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"string"},

		// values
		[]interface{}{
			"Hello",
		},
	)

	fmt.Println(hex.EncodeToString(hash))
	//dumpBlocks()
}
