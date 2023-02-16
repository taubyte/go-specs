package builders

import (
	"fmt"
	"os"
	"path"

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

type dir struct {
	wd         string
	taubyteDir string
}
type Dir interface {
	CodeSource(string) string
	TaubyteDir() string
	ConfigFile() string
	DockerDir() dockerDir
	DockerFile() string
	DefaultOptions(script string, outDir string, environment Environment) []ci.ContainerOption
	SetSourceVolume() ci.ContainerOption
	SetOutVolume(string) ci.ContainerOption
	SetBuildCommand(script string) ci.ContainerOption
	SetEnvironmentVariables() ci.ContainerOption
	String() string
}

func Wd(workDir string) (*dir, error) {
	taubyteDir := TaubyteDir
	_, err := os.Stat(path.Join(workDir, TaubyteDir))
	if err != nil {
		taubyteDir = DepreciatedTaubyteDir
		_, err := os.Stat(path.Join(workDir, taubyteDir))
		if err != nil {
			return nil, fmt.Errorf("no taubyte directory found in `%s`", workDir)
		}
	}

	return &dir{
		wd:         workDir,
		taubyteDir: taubyteDir,
	}, nil
}

type dockerDir string

type ExtraVolume struct {
	SourcePath                       string
	ContainerPath                    string
	SourceIsRelativeToBuildDirectory bool
}
