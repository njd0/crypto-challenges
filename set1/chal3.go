package set1

import (
	"crypto/m/helper/score"
	"encoding/hex"
	"fmt"
)

func DecodeAndScoreText(input string) (string, error) {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}

	// track highest ranked answer
	maxScore := -1.0
	answer := ""

  // Try all possible single-byte keys (0x00 to 0xFF)
	for i := 0; i < 256; i++ {
		result := make([]byte, len(decoded))
		for j, b := range decoded {
      result[j] = b ^ byte(i)
    }

    if score := score.ScoreText(string(result)); score > maxScore {
      maxScore = score
      answer = string(result)
    }
	}

  return answer, nil
}

// Set 1 Challenge 3
func Chal3(input string) (string, error) {
	fmt.Println("Set 1 Chal3")

  answer, err := DecodeAndScoreText(input)

  fmt.Println("answer: ", answer)
  return answer, err
}