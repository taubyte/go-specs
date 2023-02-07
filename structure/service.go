package structureSpec

type Service struct {
	Id          string
	Name        string
	Description string
	Tags        []string

	Protocol string

	// noset, this is parsed from the tags
	SmartOps []string
}

func (s Service) GetName() string {
	return s.Name
}

func (s *Service) SetId(id string) {
	s.Id = id
}
