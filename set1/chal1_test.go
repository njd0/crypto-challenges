package set1

import (
	"crypto/m/config"
	"testing"
)

func TestChal1(t *testing.T) {
	config := config.GetConfig()
	if Chal1(config.Set1Chal1.Input) != config.Set1Chal1.Output {
		t.Errorf("Not expected input")
	}
}
