package builders

import "time"

var DepreciatedTaubyteDir = false

const (
	Source          = "src"
	Output          = "out"
	DockerDir       = "Docker"
	TaubyteDir      = ".taubyte"
	TaubyteDirOld   = "taubyte"
	Dockerfile      = "Dockerfile"
	ScriptExtension = ".sh"
	ConfigFile      = "config.yaml"
)

var (
	ImageCleanInterval = 24 * time.Hour
	ImageCleanAge      = 7 * ImageCleanInterval
	DefaultTime        = time.Unix(0, 0)
)
