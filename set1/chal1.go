package set1

import (
	"encoding/hex"
	"fmt"
	"runtime"
)

var base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
// Set 1 Challenge 1
func Chal1(input string) (string){
	fmt.Println("Set1 Chal1:")

	// 1. Convert hex string to bytes
	inputBytes, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		runtime.Goexit()
	}

	var encoded string
	padding := 0
	// 2. Process the input in 3-byte chunks
	for i := 0; i < len(inputBytes); i += 3 {
		var chunk [3]byte
		copy(chunk[:], inputBytes[i:])
		chunkLength := len(inputBytes) - i

		if chunkLength < 3 {
			// if last chunk does not have 3 bytes, pad remaining
			padding = 3 - chunkLength
		}

		// 3. split 24 bits into chunks of 4 6-bit chunks
		cI1 := chunk[0] >> 2 // 8 bits -> take 6 MSB (2 rem)
		cI2 := (chunk[0] & 0x3) << 4 | chunk[1] >> 4 // 2 LSB (2 rem) && 4 MSB of 2nd chunk = 6 
		cI3 := (chunk[1] & 0x0F) << 2 | chunk[2] >> 6 // 4 LSB of 2nd chunk && 2 MSB of 3rd chunk = 6
		cI4 := chunk[2] & 0x3F// 6 LSB of 3rd chunk

		encoded += string(base64Chars[cI1])
		encoded += string(base64Chars[cI2])

		if chunkLength > 1 {
			encoded += string(base64Chars[cI3])
		}
		
		if chunkLength > 2 {
			encoded += string(base64Chars[cI4])
		}
	}
	
	for i := 0; i < padding; i++ {
		encoded += "="
	}

	fmt.Println("answer:", encoded)
  return encoded
}

