package domainSpec

import (
	"regexp"

	"github.com/taubyte/go-specs/common"
)

const PathVariable common.PathVariable = "domains"

var SpecialDomain = regexp.MustCompile(`^[^.]+\.g\.tau\.link$`)
