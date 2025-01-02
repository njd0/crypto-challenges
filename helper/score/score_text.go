package score

import (
	"crypto/m/helper/dictionary"
	"regexp"
	"strings"
)

var letterFrequencies = map[rune]float64{
	'a': 8.167, 'b': 1.492, 'c': 2.782, 'd': 4.253, 'e': 12.702,
	'f': 2.228, 'g': 2.015, 'h': 6.094, 'i': 6.966, 'j': 0.153,
	'k': 0.772, 'l': 4.025, 'm': 2.406, 'n': 6.749, 'o': 7.507,
	'p': 1.929, 'q': 0.095, 'r': 5.987, 's': 6.327, 't': 9.056,
	'u': 2.758, 'v': 0.978, 'w': 2.361, 'x': 0.150, 'y': 1.974,
	'z': 0.074, ' ': 15.000,
}

// scoreText calculates a score based on English letter frequency
func FrequencyScoreText(text string) float64 {
	score := 0.0
	for _, char := range strings.ToLower(text) {
		score += letterFrequencies[char]
	}
	return score
}

// todo make this more efficient
func DictionaryScoreText(text string) float64 {
	// regex for splitting by whitespace
	re := regexp.MustCompile(`\s+`)

	// Split the string
	parts := re.Split(text, -1)

	dictionary, err := dictionary.LoadDictionary()
	if err != nil {
		panic(err)
	}

	// score the message 
	score := 0.0
	for _, word := range parts {
		if dictionary[strings.ToLower((word))] {
			score++
		}
	}

	return score
}

// line of text is split into words and input is scored
//  split is currently on whitespace
func ScoreText(text string) float64 {
	return FrequencyScoreText(text)
}