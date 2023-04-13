package methods

import (
	"github.com/taubyte/go-specs/common"
)

func ProjectPrefix(projectId, branch, commit string) *common.TnsPath {
	return common.NewTnsPath([]string{
		common.BranchPathVariable.String(),
		branch,
		common.CommitPathVariable.String(),
		commit,
		common.ProjectPathVariable.String(),
		projectId,
	})
}
