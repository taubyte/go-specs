package builders

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"

	ci "github.com/taubyte/go-simple-container"
	"github.com/taubyte/utils/bundle"
)

func (d *dir) String() string {
	return string(*d)
}

// func (d *dir) SourceDir() string {
// 	return path.Join(d.String(), Source)
// }

func (d *dir) CodeSource(file string) string {
	return path.Join(d.String(), file)
}

// func (d *dir) OutDir() string {
// 	return path.Join(d.String(), Output)
// }

func (d *dir) TaubyteDir() string {
	return path.Join(d.String(), taubyteDir())
}

func (d *dir) ConfigFile() string {
	return path.Join(d.TaubyteDir(), ConfigFile)
}

func (d *dir) DockerDir() dockerDir {
	return dockerDir(path.Join(d.TaubyteDir(), DockerDir))
}

func (d *dir) DockerFile() string {
	return path.Join(d.DockerDir().String(), Dockerfile)
}

func (d *dir) SetSourceVolume() ci.ContainerOption {
	return ci.Volume(d.String(), "/"+Source)
}

func (d *dir) SetOutVolume(dir string) ci.ContainerOption {
	return ci.Volume(dir, "/"+Output)
}

func (d *dir) SetEnvironmentVariables() ci.ContainerOption {
	return ci.Variables(map[string]string{
		"OUT": "/" + Output,
		"SRC": "/" + Source,
	})
}

func (d *dir) SetBuildCommand(script string) ci.ContainerOption {
	return ci.Command([]string{"/bin/sh", "/" + Source + "/" + taubyteDir() + "/" + script + ScriptExtension})
}

func (d *dir) DefaultOptions(script, outDir string, environment Environment) []ci.ContainerOption {
	ops := []ci.ContainerOption{
		d.SetSourceVolume(),
		d.SetEnvironmentVariables(),
		d.SetOutVolume(outDir),
		d.SetBuildCommand(script),
	}

	if len(environment.Variables) > 0 {
		ops = append(ops, ci.Variables(environment.Variables))
	}

	return ops
}

func ExtraVolumes(wd string, volumes ...ExtraVolume) ([]ci.ContainerOption, error) {
	options := make([]ci.ContainerOption, 0, len(volumes))
	for _, volume := range volumes {
		if len(volume.SourcePath) == 0 || len(volume.ContainerPath) == 0 {
			return nil, errors.New("attempting to parse an empty volume")
		}

		if volume.SourceIsRelativeToBuildDirectory == true {
			volume.SourcePath = path.Join(wd, volume.SourcePath)
		}

		if string(volume.ContainerPath[0]) == "/" {
			options = append(options, ci.Volume(volume.SourcePath, volume.ContainerPath))
		} else {
			options = append(options, ci.Volume(volume.SourcePath, "/"+volume.ContainerPath))
		}
	}

	return options, nil
}

func (d dockerDir) String() string {
	return string(d)
}

func (d dockerDir) Stat() (fs.FileInfo, error) {
	fileInfo, err := os.Stat(string(d))
	if err != nil {
		return nil, err
	}
	if !fileInfo.IsDir() {
		return nil, fmt.Errorf("expected path `%s` to be directory, got file", d)
	}

	return fileInfo, nil
}
func (d dockerDir) Tar() ([]byte, error) {
	ops := bundle.Options{FileOptions: bundle.FileOptions{
		AccessTime: DefaultTime,
		ChangeTime: DefaultTime,
		ModTime:    DefaultTime,
	}}

	var buf bytes.Buffer
	err := bundle.Tarball(string(d), &ops, &buf)
	if err != nil {
		return nil, fmt.Errorf("tarball failed with: %s", err)
	}

	return buf.Bytes(), nil
}

/******************** Backwards Compatibility  ************************/

func taubyteDir() string {
	taubyteDir := TaubyteDir
	if DepreciatedTaubyteDir == true {
		taubyteDir = TaubyteDirOld
	}

	return taubyteDir
}

func (c *Config) HandleDepreciatedEnvironment() (environment Environment) {
	if len(c.Environment.Image) == 0 {
		return c.Enviroment
	}

	return c.Environment
}
