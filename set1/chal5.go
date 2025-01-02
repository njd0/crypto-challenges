package set1

import (
	"crypto/m/helper/encryption"
	"encoding/hex"
	"fmt"
)

func Chal5(input string, key string) (answer string, error error) {
	fmt.Println("Set 1 Chal5")

	encoded := encryption.RepeatingKey([]byte(input), []byte(key))

	answer = hex.EncodeToString(encoded)
	
	fmt.Println("answers: ", answer)
	return answer, nil
}