package encryption

// iterate through key indexes and xor each next input byte with next key byte, repeat
func RepeatingKey(input []byte, key []byte) []byte {
		result := make([]byte, len(input))
		for i, b := range input {
      result[i] = b ^ key[i % len(key)]
    }

		return result
}