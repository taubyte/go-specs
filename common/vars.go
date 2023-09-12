package common

const (
	ProjectPathVariable     PathVariable = "projects"
	ApplicationPathVariable PathVariable = "applications"
	ProjectIndexVariable    PathVariable = "project"

	/********************** Versioning VARS **********************/

	LinksPathVariable         PathVariable = "links"
	BranchPathVariable        PathVariable = "branches"
	CommitPathVariable        PathVariable = "commit"
	CurrentCommitPathVariable PathVariable = "current"
)

// TODO remove this and iterate, default branch should be gathered from a given repository
var DefaultBranch = "master"

const (
	Auth      = "auth"
	Patrick   = "patrick"
	Monkey    = "monkey"
	TNS       = "tns"
	Hoarder   = "hoarder"
	Substrate = "substrate"
	Seer      = "seer"
	Gateway   = "gateway"
)

var (
	Protocols          = []string{Auth, Patrick, Monkey, TNS, Hoarder, Substrate, Seer, Gateway}
	HTTPProtocols      = []string{Patrick, Substrate, Seer, Auth, Gateway}
	P2PStreamProtocols = []string{Seer, Auth, Patrick, TNS, Monkey, Hoarder, Substrate}
)
