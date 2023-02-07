package structureSpec

type Website struct {
	Id          string
	Name        string
	Description string
	Tags        []string
	Domains     []string
	Paths       []string
	Branch      string
	Provider    string
	RepoID      string `mapstructure:"repository-id"`
	RepoName    string `mapstructure:"repository-name"`

	// noset, this is parsed from the tags
	SmartOps []string
}

func (w Website) GetName() string {
	return w.Name
}

func (f *Website) SetId(id string) {
	f.Id = id
}
