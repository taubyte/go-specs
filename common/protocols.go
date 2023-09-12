package common

import (
	"fmt"

	"golang.org/x/exp/slices"
)

const (
	TnsProtocol  = "/tns/v1"
	SeerProtocol = "/seer/v1"
)

type ValidateOption func(string) error

func ValidateProtocols(protocols []string, ops ...ValidateOption) error {
	for _, protocol := range protocols {
		if !slices.Contains(Protocols, protocol) {
			return fmt.Errorf("`%s` is not a valid protocol", protocol)
		}

		for _, op := range ops {
			if err := op(protocol); err != nil {
				return err
			}
		}
	}

	return nil
}

func ValidateHttp() ValidateOption {
	return func(protocol string) error {
		if !slices.Contains(HTTPProtocols, protocol) {
			return fmt.Errorf("`%s` is not a http protocol", protocol)
		}
		return nil
	}
}

func ValidateP2P() ValidateOption {
	return func(protocol string) error {
		if !slices.Contains(P2PStreamProtocols, protocol) {
			return fmt.Errorf("`%s` is not a p2p stream protocol", protocol)
		}
		return nil
	}
}
