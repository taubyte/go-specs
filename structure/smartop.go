package structureSpec

type SmartOp struct {
	Id          string
	Name        string
	Description string
	Tags        []string

	Timeout uint64
	Memory  uint64
	Call    string
	Source  string

	// noset, this is parsed from the tags
	SmartOps []string
}

func (s SmartOp) GetName() string {
	return s.Name
}

func (s *SmartOp) SetId(id string) {
	s.Id = id
}
