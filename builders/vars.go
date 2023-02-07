package builders

import "time"

const (
	Source          = "src"
	Output          = "out"
	TaubyteDir      = "taubyte"
	DockerDir       = "Docker"
	Dockerfile      = "Dockerfile"
	ScriptExtension = ".sh"
	ConfigFile      = "config.yaml"
)

var (
	ImageCleanInterval = 24 * time.Hour
	ImageCleanAge      = 7 * ImageCleanInterval
	DefaultTime        = time.Unix(0, 0)
)
