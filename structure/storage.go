package structureSpec

type Storage struct {
	Id          string
	Name        string
	Description string
	Tags        []string

	Match      string
	Regex      bool `mapstructure:"useRegex"`
	Type       string
	Public     bool
	Size       uint64
	Ttl        uint64
	Versioning bool

	// noset, this is parsed from the tags
	SmartOps []string
}

func (s Storage) GetName() string {
	return s.Name
}

func (s *Storage) SetId(id string) {
	s.Id = id
}
