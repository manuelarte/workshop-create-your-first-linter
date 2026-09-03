# Workshop: Create Your First Linter

Welcome to this workshop on how to create your first linter.

## Steps

## Check AST example

In the folder [astexample](astexample) we can see a small application to
print the Abstract Syntax Tree (AST) of a Go file.

We are going to first check the AST for [example.go](astexample/testdata/src/simple/example.go).

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

### How to implement it?

Let's apply TDD, and come up with some testdata cases.

After that, we can implement the linter.

- Go through the AST and find the constants.
- Check if the constant matches the criteria:
   - Is it unexported?
   - Does it start with a lowercase letter?
   - ...
- Extra step: Propose a fix.

## Integrate it in Golangci-lint

TODO