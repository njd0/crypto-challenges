package set1

import (
	"encoding/hex"
	"fmt"
	"runtime"
)

// Set 1 Challenge 1
func Chal2(inputs ...string) (string) {
	fmt.Println("Set1 Chal2")
	buffer1 := inputs[0]
	buffer2 := inputs[1]

	bufferBytes1, bufferErr1 := hex.DecodeString(buffer1)

	if bufferErr1 != nil {
		fmt.Println("Error decoding hex string:", bufferErr1)
		runtime.Goexit()
	}

	bufferBytes2, bufferErr2 := hex.DecodeString(buffer2)

	if bufferErr2 != nil {
		fmt.Println("Error decoding hex string:", bufferErr2)
		runtime.Goexit()
	}

	if len(bufferBytes1) != len(bufferBytes2) {
		fmt.Println("Not equal buffers")
		runtime.Goexit()
	}

	var bufferLength = len(bufferBytes1)
	var answer []byte
	for i := 0; i < bufferLength; i++ {
		xor := bufferBytes1[i] ^ bufferBytes2[i]
		answer = append(answer, xor)
	}
	
  fmt.Println("answer: ", string(answer))
  return hex.EncodeToString(answer)
}

