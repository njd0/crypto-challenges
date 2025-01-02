package set1

import (
	"crypto/m/helper/file"
	"crypto/m/helper/score"
	"fmt"
)

func Chal4(input string) (string, error) {
	fmt.Println("Set 1 Chal4")
	var fileText []string

	err := file.ProcessFile(input, func(line string) {
		fileText = append(fileText, line)
	})

	if err != nil {
		panic(err)
	}

	maxScore := -1.0
	answer := ""
	for _, line := range fileText {
		decodedLine, err := DecodeAndScoreText(line)
		if err != nil {
			continue
		}

		if score := score.ScoreText(string(decodedLine)); score > maxScore {
      maxScore = score
      answer = string(decodedLine)
    }
	}

	fmt.Println("answer: ", answer)
	return answer, nil
}