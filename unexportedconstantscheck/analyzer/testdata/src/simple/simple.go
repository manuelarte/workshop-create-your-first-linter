package simple

const fooSimple = "bar" // want `unexported constant "fooSimple" should start with "_"`

const _nameSimple = "John"

const errCodeSimple = 500

const ExportedSimple = "baz"

func myFunction() {
	const localConst = "local"
}
