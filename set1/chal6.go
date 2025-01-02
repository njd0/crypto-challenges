package set1

import (
	"crypto/m/helper/file"
	"crypto/m/helper/structures"
	"fmt"
	"math"
)

func hammingDistance (input1 []byte, input2 []byte) int {
	if len(input1) != len (input2) {
		panic("Input byte lengths dont match")
	}
	
	// sum the differing bits
	// isolate each bit and compare for differences
	hammingDistance := 0
	for i:= 0; i < len(input1); i++ {
		for b := 0.0; b < 8; b++ {
			mask := byte(math.Pow(2, b))
			input1Bit := (input1[i] & mask)
			input2Bit := (input2[i] & mask)
			if input1Bit != input2Bit {
				hammingDistance++
			}
		}
	}

	return hammingDistance
}

func Chal6() (answer string, error error) {
	fmt.Println("Set 1 Chal6")

	var encodedText []byte
	file.ProcessFile("./set1/chal6.txt", func (text string) {
		encodedText = append(encodedText, text...)
	})

	heap := structures.NewSortedMinHeap(3)
	for KEYSIZE := 2; KEYSIZE < 41; KEYSIZE++ {
		// todo make programatic 
		var chunk1 []byte = make([]byte, KEYSIZE)
		var chunk2 []byte = make([]byte, KEYSIZE)
		var chunk3 []byte = make([]byte, KEYSIZE)
		var chunk4 []byte = make([]byte, KEYSIZE)

		copy(chunk1[:], encodedText[:])
		copy(chunk2[:], encodedText[KEYSIZE:])
		copy(chunk3[:], encodedText[KEYSIZE*2:])
		copy(chunk4[:], encodedText[KEYSIZE*3:])

		distance1 := hammingDistance(chunk1, chunk2) / KEYSIZE
		distance2 := hammingDistance(chunk3, chunk4) / KEYSIZE
		normalizedDistance := (distance1 + distance2) / 2

		heap.Add((normalizedDistance))
	}

	

	fmt.Println(heap.Data)
	return answer, nil
}