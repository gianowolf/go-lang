# Learning Go - Notes

"Learning Go: An Idiomatic Approach to Real-World Go Programming" J. Bonder, O'Reilly (2022)

## 01 Setting Up Go Environment

- Go Tools: https://golang.org/dl

```sh
go version
```

### Workspace

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### Go Command

- ```go``` tool: compiler, code formatter, linter, dependency manager, test runner, and more.

#### Go Run - Go Build

- Takes a Go file, list of Go files or a package name.
  
```sh
# hello.go
package main

import "fmt"

func main(){
    fmt.Println("Hello, world!")
}
```


```sh
go run hello.go
```