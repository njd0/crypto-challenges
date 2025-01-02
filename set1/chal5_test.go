package set1

import (
	"crypto/m/config"
	"testing"
)

func TestChal5(t *testing.T) {
	config := config.GetConfig()
	answer, err := Chal5(config.Set1Chal5.Inputs.Text, config.Set1Chal5.Inputs.Key)
	if err != nil {
		t.Errorf("Error in Chal5 routine")
	}

	if answer != config.Set1Chal5.Output {
		t.Errorf("Not expected input")
	}
}
