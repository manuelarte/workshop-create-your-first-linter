package main

/* Look all the different ways I can declare constants in Go */

// This Go spec is a *ast.GenDecl with Tok of token.CONST.
// There is one *ast.ValueSpec.
const foo = "foo"

// This Go spec is a *ast.GenDecl with Tok of token.CONST.
// There are two *ast.ValueSpec.
const (
	foo1 = "foo1"
	foo2 = "foo2"
)

// This Go spec is a *ast.GenDecl with Tok of token.CONST.
// There is one *ast.ValueSpec with two names.
const foo3, foo4 = "foo3", "foo4"

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
