package structureSpec

type Library struct {
	Id          string
	Name        string
	Description string
	Tags        []string

	Path     string
	Branch   string
	Provider string
	RepoID   string `mapstructure:"repository-id"`
	RepoName string `mapstructure:"repository-name"`

	// noset, this is parsed from the tags
	SmartOps []string
}

func (l Library) GetName() string {
	return l.Name
}

func (l *Library) SetId(id string) {
	l.Id = id
}
