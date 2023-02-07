package domainSpec

import "testing"

func TestValidate(t *testing.T) {
	err := ValidateDNS("test-id", "00test-id.g.tau.link", true)
	if err != nil {
		t.Error(err)
		return
	}
}
