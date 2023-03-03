package structureSpec

type Database struct {
	Id          string
	Name        string
	Description string
	Tags        []string

	Match string
	Regex bool `mapstructure:"useRegex"`
	Local bool
	Key   string
	Min   int
	Max   int
	Size  uint64

	// noset, this is parsed from the tags
	SmartOps []string
}

func (d Database) GetName() string {
	return d.Name
}

func (d *Database) SetId(id string) {
	d.Id = id
}
