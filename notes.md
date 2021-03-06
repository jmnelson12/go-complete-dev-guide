# Notes

## Common Go CLI Commands

- `go build` - compiles a bunch of `go` source code files
- `go run` - compiles and executes one or two files
- `go fmt` - formats all of the code in each file in the current directory
- `go install` - compiles and 'installs' a package
- `go get` - downloads the raw source code of someone elses package
- `go test` - runs any tests associated with the current project

## Packages

https://golang.org/pkg

## Playground

https://play.golang.org/

## Some Basic Data Types

- bool
- string
- int
- float64
- array - fixed length of typed items
- slice - array that can grow or shrink
- map

## Questions and Answers

How do we run the code in our project?

- `go run <filename>.go`

What does `package main` mean?

- `package main`

  - Creates an executable package that is compiled then executed
  - **Must** have a func called `main`
  - `main` is special

- `package calculator` or `package uploader`
  - creates a reusable package that is used as a dependency (helper code)

How is the `main.go` file organized?

```go
# Package declaration
package main

# Import other packages that we need
import "fmt"

# The code
func main() {
  fmt.Println("Hello")
}
```

---

## Value Types vs Reference Types

| Value Types (need pointers) | Reference Types (don't need pointers) |
| --------------------------- | ------------------------------------- |
| int                         | slices                                |
| float                       | maps                                  |
| string                      | channels                              |
| bool                        | pointers                              |
| structs                     | function                              |

---

## Map vs Struct

| Map                                                 | Struct                                                       |
| --------------------------------------------------- | ------------------------------------------------------------ |
| keys must be same type                              | values can be of different type                              |
| values must be same type                            |                                                              |
| can be iterated over                                | keys can't be indexed                                        |
| use to represent a collection of related properties | need to know all fields at compile time                      |
| reference type                                      | value type                                                   |
| Dont need to know all the keys at compile time      | use to represent a "thing" with lots of different properties |
