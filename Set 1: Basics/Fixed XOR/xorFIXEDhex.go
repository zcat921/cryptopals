//XOR'ing two fixed length hex values
//Usage: go run xorFIXEDhex.go <hex_value 1> <hex_value 2>
package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

func main() {
	inp1 := os.Args[1]
	inp2 := os.Args[2]
	dh1, err1 := hex.DecodeString(inp1)
	dh2, err2 := hex.DecodeString(inp2)

	if err1 != nil {
		panic(err1)
	} else if err2 != nil {
		panic(err2)
	}

	fmt.Println("Decoded Hex 1 :", string(dh1), "=", dh1)
	fmt.Println("Decoded Hex 2 :", string(dh2), "=", dh2)

	xored, _ := XORDecrypt(dh1, dh2)
	hex_out := hex.EncodeToString([]byte(xored))
	fmt.Println("Hex Converted text :", string(hex_out))

}

func XORDecrypt(val1, val2 []byte) ([]byte, error) {
	if len(val1) != len(val2) {
		return nil, errors.New("Inputs not have equal length")
	}
	out := make([]byte, len(val1))
	for i := 0; i < len(val1); i++ {
		out[i] = val1[i] ^ val2[i]
	}
	return out, nil
}
