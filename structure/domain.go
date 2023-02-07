package structureSpec

type Domain struct {
	Id          string
	Name        string
	Description string
	Tags        []string

	Fqdn     string
	CertType string `mapstructure:"cert-type"`
	CertFile string `mapstructure:"cert-file"`
	KeyFile  string `mapstructure:"key-file"`

	// noset, this is parsed from the tags
	SmartOps []string
}

func (d Domain) GetName() string {
	return d.Name
}

func (d *Domain) SetId(id string) {
	d.Id = id
}
