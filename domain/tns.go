package domainSpec

import (
	"errors"
	"strings"

	"github.com/taubyte/go-specs/common"
	"github.com/taubyte/go-specs/methods"

	slices "github.com/taubyte/utils/slices/string"
)

func Tns() *tnsHelper {
	return &tnsHelper{}
}

func (t *tnsHelper) BasicPath(fqdn string) (*common.TnsPath, error) {
	if fqdn == "" {
		return nil, errors.New("Fqdn is empty")
	}

	array_to_reverse := strings.Split(fqdn, ".")
	reversed := slices.ReverseArray(array_to_reverse)

	return common.NewTnsPath(append([]string{string(PathVariable)}, reversed...)), nil
}

func (t *tnsHelper) IndexValue(branch, projectId, appId, domId string) (*common.TnsPath, error) {
	return methods.IndexValue(branch, projectId, appId, domId, PathVariable)
}
