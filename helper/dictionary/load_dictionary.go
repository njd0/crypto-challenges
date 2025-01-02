package dictionary

import (
	"crypto/m/helper/file"
	"strings"
)

// LoadDictionary loads words from a dictionary file into a map
func LoadDictionary() (map[string]bool, error) {
	words := make(map[string]bool)

	err := file.ProcessFile("/usr/share/dict/words", func (line string) {
		word := strings.ToLower(line)
		words[word] = true
	})

	return words, err
}