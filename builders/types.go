package builders

import (
	ci "github.com/taubyte/go-simple-container"
)

type Environment struct {
	Image     string
	Variables map[string]string
}

type Config struct {
	Version     string
	Environment Environment
	Workflow    []string

	// TODO: Repo Fixer should remove all of these cases for websites and libraries
	Enviroment Environment
}

type dir string
type Dir interface {
	SourceDir() string
	CodeSource(string) string
	OutDir() string
	TaubyteDir() string
	ConfigFile() string
	DockerDir() dockerDir
	DockerFile() string
	DefaultOptions(script string, environment Environment) []ci.ContainerOption
	SetSourceVolume() ci.ContainerOption
	SetOutVolume() ci.ContainerOption
	SetBuildCommand(script string) ci.ContainerOption
	SetEnvironmentVariables() ci.ContainerOption
	String() string
}

func Wd(workDir string) *dir {
	wd := dir(workDir)
	return &wd
}

type dockerDir string

type ExtraVolume struct {
	SourcePath                       string
	ContainerPath                    string
	SourceIsRelativeToBuildDirectory bool
}
