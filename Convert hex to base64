//This programme convert Hex value to Base64
package main
import (
	"fmt"
	"os"
	"encoding/hex"
	"encoding/base64"
)

func main() {
	str := os.Args[1]
	bs,err:=hex.DecodeString(str)
	if err!=nil {
		panic(err)
	}
	fmt.Println("Decoded Hex :",string(bs))

	b64:=base64.StdEncoding.EncodeToString([]byte(bs))
	fmt.Println("base64 Encoded text :",string(b64))
}
