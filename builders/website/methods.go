package website

import (
	"path"

	ci "github.com/taubyte/go-simple-container"
	"github.com/taubyte/go-specs/builders"
)

func Wd(workDir string) (dir Dir, err error) {
	wd, err := builders.Wd(workDir)
	if err != nil {
		return
	}

	return Dir{wd}, nil
}

func (d Dir) BuildZip() string {
	return path.Join(d.String(), ZipFile)
}

func (d Dir) SetWorkDir() ci.ContainerOption {
	return ci.WorkDir("/" + builders.Source)
}
