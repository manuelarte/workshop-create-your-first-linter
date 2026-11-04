package simple

const (
	fooGroups      = "bar" // want `unexported constant "fooGroups" should start with "_"`
	_nameGroups    = "John"
	errCodeGroups  = 500
	ExportedGroups = "baz"
)
