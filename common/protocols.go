package common

import (
	"fmt"

	slices "github.com/taubyte/utils/slices/string"
)

const (
	TnsProtocol  = "/tns/v1"
	SeerProtocol = "/seer/v1"
)

func ValidateProtocols(protocols []string) error {
	for _, protocol := range protocols {
		if !slices.Contains(Protocols, protocol) {
			return fmt.Errorf("`%s` is not a valid protocol", protocol)
		}
	}

	return nil
}
