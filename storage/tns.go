package storageSpec

import (
	"github.com/taubyte/go-specs/common"
	"github.com/taubyte/go-specs/methods"
)

func Tns() *tnsHelper {
	return &tnsHelper{}
}

func (t *tnsHelper) BasicPath(branch, commit, projectId, appId, storeId string) (*common.TnsPath, error) {
	return methods.GetBasicTNSKey(branch, commit, projectId, appId, storeId, PathVariable)
}

func (t *tnsHelper) IndexValue(branch, projectId, appId, storId string) (*common.TnsPath, error) {
	return methods.IndexValue(branch, projectId, appId, storId, PathVariable)
}

func (t *tnsHelper) IndexPath(projectId, appId, name string) *common.TnsPath {
	if len(appId) > 0 {
		return common.NewTnsPath([]string{
			common.ProjectPathVariable.String(),
			projectId,
			common.ApplicationPathVariable.String(),
			appId,
			name,
		})
	}

	return common.NewTnsPath([]string{common.ProjectPathVariable.String(), projectId, name})
}
