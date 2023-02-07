package structureSpec

type Messaging struct {
	Id          string
	Name        string
	Description string
	Tags        []string

	Local     bool
	Match     string
	Regex     bool
	MQTT      bool
	WebSocket bool

	// noset, this is parsed from the tags
	SmartOps []string
}

func (m Messaging) GetName() string {
	return m.Name
}

func (m *Messaging) SetId(id string) {
	m.Id = id
}
