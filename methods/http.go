package methods

import (
	"errors"
	"strings"

	"github.com/taubyte/go-specs/common"
	slices "github.com/taubyte/utils/slices/string"
)

func HttpPath(fqdn string, resourceType common.PathVariable) (*common.TnsPath, error) {
	if fqdn == "" {
		return nil, errors.New("Fqdn is empty")
	}

	array_to_reverse := strings.Split(fqdn, ".")
	reversed := slices.ReverseArray(array_to_reverse)

	return common.NewTnsPath(append([]string{"http", string(resourceType)}, reversed...)), nil
}
