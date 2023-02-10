package website

import (
	"path"

	ci "github.com/taubyte/go-simple-container"
	"github.com/taubyte/go-specs/builders"
)

func Wd(workDir string) Dir {
	wd := builders.Wd(workDir)
	return Dir{wd}
}

func (d Dir) BuildZip() string {
	return path.Join(d.String(), ZipFile)
}

func (d Dir) DefaultOptions(script string, environment builders.Environment) []ci.ContainerOption {
	ops := d.Dir.DefaultOptions(script, environment)
	ops = append(ops, d.SetWorkDir())
	return ops
}

func (d Dir) SetWorkDir() ci.ContainerOption {
	return ci.WorkDir("/" + builders.Source)
}
