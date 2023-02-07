package website

import (
	"path"

	"github.com/taubyte/go-specs/builders"
	ci "github.com/taubyte/go-simple-container"
)

func Wd(workDir string) Dir {
	wd := builders.Wd(workDir)
	return Dir{wd}
}

func (d Dir) DockerFile() string {
	return path.Join(d.DockerDir(), builders.Dockerfile)
}

func (d Dir) BuildZip() string {
	return path.Join(d.String(), ZipFile)
}

func (d Dir) DefaultOptions(script string) []ci.ContainerOption {
	return []ci.ContainerOption{
		d.SetSourceVolume(),
		d.SetEnvironmentVariables(),
		d.SetOutVolume(),
		d.SetWorkDir(),
		d.SetBuildCommand(script),
	}
}

func (d Dir) SetEnvironmentVariables() ci.ContainerOption {
	return ci.Variables(map[string]string{
		"OUT": "/" + builders.Output,
		"SRC": "/" + builders.Source,
	})
}

func (d Dir) SetWorkDir() ci.ContainerOption {
	return ci.WorkDir("/" + builders.Source)
}
