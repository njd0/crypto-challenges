package set1

import (
	"crypto/m/config"
	"testing"
)

func TestChal2(t *testing.T) {
	if Chal2(config.GetConfig().Set1Chal2.Inputs...) != config.GetConfig().Set1Chal2.Output {
		t.Errorf("Not expected input")
	}
}
