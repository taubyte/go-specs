package structureSpec

import "github.com/taubyte/go-specs/common"

type Simple struct{}

func (Simple) GetName() string {
	return ""
}

func (Simple) SetId(string) {}

func (Simple) BasicPath(branch, commit, project, app string) (*common.TnsPath, error) {
	return nil, nil
}

func (Simple) GetId() string {
	return ""
}
