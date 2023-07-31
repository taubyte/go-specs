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

var Protocols = []string{"auth", "patrick", "monkey", "tns", "hoarder", "substrate", "seer"}
var HttpProtocols = []string{"patrick", "substrate", "seer", "auth"}
