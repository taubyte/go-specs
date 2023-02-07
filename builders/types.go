package builders

import ci "github.com/taubyte/go-simple-container"

type Environment struct {
	Image     string
	Variables map[string]string
}

type Config struct {
	Version     string
	Environment Environment
	Workflow    []string
}

type dir string

type Dir interface {
	SourceDir() string
	OutDir() string
	TaubyteDir() string
	ConfigFile() string
	DockerDir() string
	SetSourceVolume() ci.ContainerOption
	SetOutVolume() ci.ContainerOption
	SetBuildCommand(script string) ci.ContainerOption
	String() string
}

func Wd(workDir string) *dir {
	wd := dir(workDir)
	return &wd
}
