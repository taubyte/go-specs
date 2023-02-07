package structureSpec

type Function struct {
	Id          string
	Name        string
	Description string
	Tags        []string

	Type     string
	Timeout  uint64
	Memory   uint64
	Call     string
	Source   string
	Domains  []string
	Method   string
	Paths    []string
	Secure   bool
	Command  string
	Channel  string
	Local    bool
	Protocol string `mapstructure:"service"`
	SmartOps []string
}

func (f Function) GetName() string {
	return f.Name
}

func (f *Function) SetId(id string) {
	f.Id = id
}
