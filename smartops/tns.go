package smartOpSpec

import (
	"github.com/taubyte/go-specs/common"
	"github.com/taubyte/go-specs/methods"
)

func Tns() *tnsHelper {
	return &tnsHelper{}
}

func (t *tnsHelper) BasicPath(branch, commit, projectId, appId, smartId string) (*common.TnsPath, error) {
	return methods.GetBasicTNSKey(branch, commit, projectId, appId, smartId, PathVariable)
}

func (t *tnsHelper) IndexValue(branch, projectId, appId, smartId string) (*common.TnsPath, error) {
	return methods.IndexValue(branch, projectId, appId, smartId, PathVariable)
}

func (t *tnsHelper) WasmModulePath(projectId, appId, resourceName string) (*common.TnsPath, error) {
	return methods.WasmModulePath(projectId, appId, resourceName, PathVariable)
}

func ModuleName(name string) string {
	return PathVariable.String() + "/" + name
}
