package builders

import (
	"errors"
	"os"
	"path"

	ci "github.com/taubyte/go-simple-container"
)

func (d *dir) String() string {
	return string(*d)
}

func (d *dir) SourceDir() string {
	return path.Join(d.String(), Source)
}

func (d *dir) OutDir() string {
	return path.Join(d.String(), Output)
}

func (d *dir) TaubyteDir() string {
	return path.Join(d.SourceDir(), TaubyteDir)
}

func (d *dir) ConfigFile() string {
	return path.Join(d.TaubyteDir(), ConfigFile)
}

func (d *dir) DockerDir() string {
	return path.Join(d.TaubyteDir(), DockerDir)
}

func (d *dir) SetSourceVolume() ci.ContainerOption {
	return ci.Volume(d.SourceDir(), "/"+Source)
}

func (d *dir) SetOutVolume() ci.ContainerOption {
	outDir := d.OutDir()
	_, err := os.Stat(outDir)
	if err != nil {
		os.Mkdir(outDir, 0755)
	}
	return ci.Volume(outDir, "/"+Output)
}

func (d *dir) SetBuildCommand(script string) ci.ContainerOption {
	return ci.Command([]string{"/bin/sh", "/" + Source + "/" + TaubyteDir + "/" + script + ScriptExtension})
}

type ExtraVolume struct {
	SourcePath                       string
	ContainerPath                    string
	SourceIsRelativeToBuildDirectory bool
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
