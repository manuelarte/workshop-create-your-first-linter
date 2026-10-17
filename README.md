# Workshop: Create Your First Linter

Welcome to this workshop on how to create your first linter (a.k.a. analyzer) in Go.

The package to use is [golang/x/tools][golang/x/tools].

## Prerequisites

- Basic knowledge of Go programming.
- Git installed on your system.
- Go compiler version 1.26 or newer.

As an alternative, you can use GitHub workspaces to run the workshop in a browser
without installing anything on your system.

## Overview

This workshop contains some explanation of how a linter works, how the AST is built and
some small exercises before we actually build our first linter.

## Steps

## Check AST example

Every compiler/interpreter parses source code into an Abstract Syntax Tree (AST), and linters work on
top of that tree. This is not a Go-only concept, so in the folder [astexample](astexample) we have
several small projects, one per language, that print the AST of a small source file:

| Language | Project | Command |
| --- | --- | --- |
| Go | [astexample/go](astexample/go) | `go test -v` |
| Rust | [astexample/rust](astexample/rust) | `just ast` |
| Python | [astexample/python](astexample/python) | `just ast` |

Feel free to explore all of them to see how similar code is represented in each language. From now
on, **we are going to focus on Go**, as that is the language we will use to build our linter.

We are going to first check the AST for
[example.go](astexample/go/testdata/src/simple/example.go).
This "linter" only prints the AST tree with:

```go
ast.Print(pass.Fset, file)
```

If we run the tests inside [astexample/go](astexample/go) with:

```bash
go test -v
```

We can see the AST tree.

### Exercise

Modify the Go example file and see the AST representation of other [go specs][go-specs].

Put special emphasis on checking how AST represents:

- Global and local: [Constants](https://go.dev/ref/spec#Constants).
- Global and local: [Variables](https://go.dev/ref/spec#Variables).

## Define our linter

The linter we are going to implement is:

[Prefix Unexported Globals With _](https://github.com/uber-go/guide/blob/master/style.md#prefix-unexported-globals-with-_)

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
// foo.go

const (
  defaultPort = 8080
  defaultUser = "user"
)

// bar.go

func Bar() {
  defaultPort := 9090
  ...
  fmt.Println("Default port", defaultPort)

  // We will not see a compile error if the first line of
  // Bar() is deleted.
}
```

</td><td>

```go
// foo.go

const (
  _defaultPort = 8080
  _defaultUser = "user"
)
```

</td></tr>
</tbody></table>

We are going to implement the linter in the folder [unexportedconstantscheck](./unexportedconstantscheck).

### Define Test Cases

Let's apply [Test-Driven Development (TDD)][TDD] and come up with some test cases and the expected output.

#### Exercise

Inside the folder `unexportedconstantscheck`, create as many testcases (in the same or different files) as 
you can think of in [analyzer/testdata/src](unexportedconstantscheck/analyzer/testdata/src).

The documentation in [analysistest][analysistest] contains how to define the test cases, but as a summary:

```go
package main

const _defaultPort = 8080 // want `my expected message`
```

### Implementation

After defining the tests, we implement the linter.

The workflow is the following:

- Go through the AST and find the global constants.
- Check if the constant matches the criteria:
   - Is it unexported?
   - Does it start with a lowercase letter?
   - ...
- Extra step: Propose a fix.

#### Exercise: Find the global constants

Inside the `run` function of the analyzer, we can use get the files
and traverse the AST:

```go
for _, file := range pass.Files {
    for _, decl := range file.Decls {
	    // TODO
	}
}
```

#### Exercise: Report the cases

Once we find something to report, we can diagnosis using `pass.Reportf`
to report the issue.

```go
analysis.Diagnostic{
    Pos:     node.Pos(),
	End:     node.End(),
	Message: msg,
}
```

### Verification

Verify all the tests pass and the error message is the expected one.

#### Extra Exercise: Suggest a fix

The analyzer tool also provides the possibility to
apply a fix to the code. This is done by providing a `SuggestedFix` in the `Diagnostic` struct.

```go
analysis.Diagnostic{
    Pos:     i.Pos(),
	End:     i.End(),
	Message: msg,
	SuggestedFixes: []analysis.SuggestedFix{
	    {
		    Message: msg,
			TextEdits: []analysis.TextEdit{
			    {
	                Pos: i.Pos(),
					End: i.End(),
					// Can someone see the problem of fixing it?
					NewText: fmt.Appendf(nil, "_%s", i.Name),
				},
			},
		},
	},
}
````

Question: Do you see any issue when suggesting a fix?

## Integrate it in Golangci-lint

We are going to add it to golangci-lint as a module 
[plugin][golangci-plugin].

### Exercise

- Add the necessary library and the code needed in `unexportedconstantscheck`.

Inside `custompluginexample`:

- Create `custom-gcl.yml` file with your linter configuration.
- Add to `.golangci.yml` the custom linter settings with your new linter.
- Create a custom golangci-lint binary and check that the linter is reporting an issuee.

```bash
golangci-lint custom -vv
./custom-gcl run --fix ./...
```

[analysistest]: https://pkg.go.dev/golang.org/x/tools/go/analysis/analysistest
[golang/x/tools]: https://pkg.go.dev/golang.org/x/tools
[golangci-plugin]: https://golangci-lint.run/docs/plugins/module-plugins/
[go-specs]: https://go.dev/ref/spec
[TDD]: https://en.wikipedia.org/wiki/Test-driven_development
