package domainSpec

import (
	"regexp"

	"github.com/taubyte/go-specs/common"
)

const PathVariable common.PathVariable = "domains"

var WhiteListedDomains = []string{"nodes.taubyte.com", "nodes.taubyte.network"}

var TaubyteServiceDomain = regexp.MustCompile(`^[^.]+\.taubyte\.com$`)
var SpecialDomain = regexp.MustCompile(`^[^.]+\.g\.tau\.link$`)
var TaubyteHooksDomain = regexp.MustCompile(`hooks.git.taubyte.com`)
