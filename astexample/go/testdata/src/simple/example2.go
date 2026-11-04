package simple

// This Go spec is a *ast.GenDecl with Tok of token.CONST.
// There are two *ast.ValueSpec.
const (
	foo              = "bar"
	ExportedConstant = "baz"
)

// This Go spec is a *ast.GenDecl with Tok of token.TYPE.
// There are four *ast.TypeSpec, but different types of *ast.TypeSpec.Type.
type (
	a struct{}
	b struct{}
	c interface {
		Hello() string
	}
	d any
)
