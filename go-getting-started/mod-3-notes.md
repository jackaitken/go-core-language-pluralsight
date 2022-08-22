# Module 3 Notes

We're going to start this module by setting up our local dev environment. Then we'll talk about using the `go` command.

## The `go` command

If we type `go` into our command line, we'll get a long output that looks like this:

```bash
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        work        workspace maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildconstraint build constraints
        buildmode       build modes
        c               calling between Go and C
        cache           build and test caching
        environment     environment variables
        filetype        file types
        go.mod          the go.mod file
        gopath          GOPATH environment variable
        gopath-get      legacy GOPATH go get
        goproxy         module proxy protocol
        importpath      import path syntax
        modules         modules, module versions, and more
        module-get      module-aware go get
        module-auth     module authentication using go.sum
        packages        package lists and patterns
        private         configuration for downloading non-public code
        testflag        testing flags
        testfunc        testing functions
        vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.
```

The last line here is very helpful as we can simply type `go help` followed by whatever topic we want to see more information on.

One such topic that we can explore is the `go doc` command. Here we can get documentation on anything in the go standard library. For example, if we want more information on the `json` package we can type `go doc json` and we'll get this:

```bash
package json // import "encoding/json"

Package json implements encoding and decoding of JSON as defined in RFC
7159. The mapping between JSON and Go values is described in the
documentation for the Marshal and Unmarshal functions.

See "JSON and Go" for an introduction to this package:
https://golang.org/doc/articles/json_and_go.html

func Compact(dst *bytes.Buffer, src []byte) error
func HTMLEscape(dst *bytes.Buffer, src []byte)
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
func Marshal(v any) ([]byte, error)
func MarshalIndent(v any, prefix, indent string) ([]byte, error)
func Unmarshal(data []byte, v any) error
func Valid(data []byte) bool
type Decoder struct{ ... }
    func NewDecoder(r io.Reader) *Decoder
type Delim rune
type Encoder struct{ ... }
    func NewEncoder(w io.Writer) *Encoder
type InvalidUTF8Error struct{ ... }
type InvalidUnmarshalError struct{ ... }
type Marshaler interface{ ... }
type MarshalerError struct{ ... }
type Number string
type RawMessage []byte
type SyntaxError struct{ ... }
type Token any
type UnmarshalFieldError struct{ ... }
type UnmarshalTypeError struct{ ... }
type Unmarshaler interface{ ... }
type UnsupportedTypeError struct{ ... }
type UnsupportedValueError struct{ ... }
```

This gives us information about everything that is defined in the json package. If we wanted more information on the `Decoder` object, we could type `go doc json.Decoder`:

```bash
package json // import "encoding/json"

type Decoder struct {
        // Has unexported fields.
}
    A Decoder reads and decodes JSON values from an input stream.

func NewDecoder(r io.Reader) *Decoder
func (dec *Decoder) Buffered() io.Reader
func (dec *Decoder) Decode(v any) error
func (dec *Decoder) DisallowUnknownFields()
func (dec *Decoder) InputOffset() int64
func (dec *Decoder) More() bool
func (dec *Decoder) Token() (Token, error)
func (dec *Decoder) UseNumber()
```

If we want more information about the `Decode` method we could type `go doc json.Decoder.Decode`:

```bash
package json // import "encoding/json"

func (dec *Decoder) Decode(v any) error
    Decode reads the next JSON-encoded value from its input and stores it in the
    value pointed to by v.

    See the documentation for Unmarshal for details about the conversion of JSON
    into a Go value.
```

So this is a very useful way to get documentation quickly and easily.

## First application

First let's create a go module. At the root level of our directory, we're going to run the `go mod init` command. There is a legacy way to do something similar called Go workspaces, but this is not the suggested way of doing this now. Modules are the official way to write Go source code.

**A Go module is simply a directory in our harddrive that has a go.mod file**. 

The `go mod init` command takes one more argument and that is the name of the module itself. 