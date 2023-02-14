package domainSpec

import (
	"context"
	"fmt"
	"strings"

	"github.com/ipfs/go-cid"
	dv "github.com/taubyte/domain-validation"
)

// TODO move to github.com/taubyte/domain-validation
func ValidateDNS(_project, host string, dev bool, options ...dv.Option) error {
	if SpecialDomain.MatchString(host) == true {
		if dev == true {
			// For dev mode Pad project string with 0s to be at least 8 characters, otherwise the check below will panic
			if len(_project) < 8 {
				_project = fmt.Sprintf("%08s", _project)
			}
		}

		// Confirm host contains last 8 of project id
		if strings.Contains(host, strings.ToLower(_project[len(_project)-8:])) == false {
			return fmt.Errorf("generated fqdn `%s` does not contain last 8 of project id %s", host, _project)
		}

		return nil
	} else if dev == false {
		// Validate project CID
		project, err := cid.Decode(_project)
		if err != nil {
			return fmt.Errorf("decoding cid `%s` failed with: %s", _project, err)
		}

		// Check if domain is registered
		err = dv.FromDNS(context.Background(), &project, host, options...)
		// If not
		if err != nil {
			// maybe the request has a port number, trim it off
			host = ExtractHost(host)

			// then try again
			err = dv.FromDNS(context.Background(), &project, host, options...)
			if err != nil {
				return fmt.Errorf("verifying DNS of `%s` after splitting port from host failed with: %s", host, err)
			}
		}
	}

	return nil
}

func ExtractHost(hostAndPort string) string {
	return strings.ToLower(strings.Split(hostAndPort, ":")[0])
}
