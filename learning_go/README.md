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

- ```go run```:
  - no binary saved there
  1. Compile go file into a temporary binary
  2. Binary is deleted after the execution
- ```go build```:
  - Creates an executable called ```hello```.

### Formatting

- ```go fmt``` automatically reformats the code to match the standard format.
- ```goimports``` improved version of fmt
  - cleans up import statements

## 2. Primitive Types and Declarations

### Built-int Types

#### Zero Value

Go assigns a default _zero value_ to any variable that is declared but not assigned a value.

#### Literals

A _Literal_ in Go refers to writing out a number, character or string. There are four commmon kinds of literals that you'll find in Go programs.

- **Integer Literals**: Secuense of numbers (normally base ten)
  - 0b binary
  - 0o octal
  - 0x hexadecimal
  - 000_000_000 longer integer literals

- **Floating point literals** have a decimal points to indicate the fractional portion. 
  - 0.32
  - 5.03e23 the can also have exponent specified with the leter _e_ (for hexa, use _p_ insthead)

- **Rune Literals** represent characters surrounded by single quotes: ''
  - double and single quotes are _NOT_ interchangeable.
    - Unicode: 'a'
    - 8-bit octal: '\141'
    - 8 and 16-bit hexa: '\x61' and '\u0061' (used for POSIX permission flag values such 0o777)
    - 32-bit Unicode '\U0000061'
    - backslash escaped: \t \n \' \"

- **String Literals**
  - double quotes ""
  - Backquotes ``: _raw string literal_ is used to include double quotes, backslashes, newlines, etc.

```txt
"Greetings and\n \"Salutations\""

`Greetings and 
"Salutations"`
```

Literals are __untyped__ because Go avoid forcing a type until the developer specifies one.

Go uses the _default type_ for a literal

### Booleans

__bool__ can have one of two values: ```true``` or   ```false```

```go
var flag bool /* No value assigned: set to false */
var isAwesome = true
```

### Numeric Types

#### Integer Types

```
uint8, byte, rune
uint16
uint32
uint64
int8
int16
int32
int64, int
```

- Zero value for integer types is 0.
- byte is an alias for uint8
- int is int64 on most 64-bit CPUs
  
##### Integer Operators

- Arithmetic: +, -, *, /, %. Can be combined with =
- Comparation: ==, !=, >=, <=.
- Bit-manipulation: &, |, ^, &^. Can be combined with =

#### Floating Point types

- float32
- float64, float

>> A floating point number cannot represent a decimal value exactly. Do not use them to represent money or any other value that must have an exact decimal representation.

#### Complex Types

- complex64 
- complex128

Uses float32 and float64 pairs to represent real and imaginary part.

>> For numerical computing applications in Go, Gonum package takes advantage of complex numbers and provides useful libraries for linear agebra, matrices, integration and statistics.

### Strings and Runes

- Go include strings as built-in type.
- zero value for a string is the empty stirng
- Go supports unicode
- strings can be compared using >, <=, >=, !=, ==, <.
- Strings are immutable. You can reassign the value of a string variable, but cannot change the value of the string that is assigned to it.

The rune type:

- rune is an alias for the int32 type
- Use the rune type to refer to a character, not the int32.

### Explicit type conversion

Go does not allow automatic type promotion between variables. You must use type conversion when variable types do not match.

```go
var x int = 10;
var y float64 = 30.2;
var z float64 = float64(x) + y
var d int = x + int(y)
fmt.Println(z, d) // >> 40.2 40
```

>> You cannot treat another Go type as a boolean. No ther type can be converted to a bool, implicitly or explicitly

### var declarations

- Go hsa a lot of ways to declare variables
- each declaration style communicates something about how the variable is used

```go
/* var keyword, explicit type, assignment */
var x int = 10

/* if type of the = is the expected, you can leave off the type */
var x = 10

/* assign the zero value, keeping the type and drop the = */
var x int // ZERO VALUE

/* Multiple variables declaration */
var x, y int = 10 // Same type
var x, y int // Zero value
var x, y = 10, "HELLO" // DIFFERENT TYPES

/* MULTIPLE VARIABLES AT ONCE */
var (
    x int
    y = 20
    z int = 30
    d, e = 40, "hello"
    f, g string
)

/* Short declaration format:
    := replace var keyword
    uses type inference */

var x = 10
x := 10

var x, y = 10, "hello"
x, y := 10 "hello"
```

- The ```:=``` opeator also allows to assign values to existing variables.
- ```:=``` it's illegal outside a function (package level)
- ```:=``` it is the most used declaration at funciton level

```go
x := 10
x, y := 30, "hello"
```

- AVOID ```:=```: 
  - Initializing a variable to its zero value
    - use ```var x int``
  - When the literal is not the type for the variable
    - use ```var x byte = 20``` 
  - You should rarely declare variables outside of functions

##### const

```go
const x int64 = 10
```

```const``` in Go is very limited. Constants in Go are a way to give names to literals. They can only hold values that the compiler can figure out at __compile time__

- Contants can be typed or untyped
- Untyped constant works like a literal it has no type of its own, but dows a default type that is used when no other type can be inferred. A typed constant can only be directly assigned to a variable of that type

#### Unused Variables

- Go requires that every declared local variable must be read. It is a compile-time error to declare a local variable and not read its value.
- Go do not catch unused assigments: use ```golangci-lint run```

## 3. Composite Types

### Arrays

- Arrays are rarely used directly in Go
- All the elements in the array must be of the type that is specified
- First, you specify the size of the array and the type of the elemnts in the arrray 

```go
var x [3]int // Zero values 0 0 0 

var x = [12]int{1, 5:4, 6, 10: 100, 15} // 1 0 0 0 0 4 6 0 0 0 0 100 15

var x = [...]int{10, 20, 30} // 10 20 30
```

Arrays in Go are unusual because Go considers the size of the array to be part of the type of the array. This also means that you cannot use a variable to specify the size of an array because types must be resolved at compile time, not at runtime

>> You can't use type conversion to conver arrays of different sizes to identical types. Because you can't convert arrays of different sizes into each other.

### Slices

Most of the times, a slice is what you should use to a data structure that holds a sequence of values. What makes slices so useful is that the length is not part of the type of a slice. this removes the limiattions of arrays. We can write a single function that processes slices of any size. And we an grow slices as needed

Working with silices looks quite a bit like working with arrays, but there are subtle differences. The first thing to notice is that we don't specify the size of the slice when we declare it

```go
var x = []int{10, 20, 30}
// [n] or [...] makes an array
// [] makes an slice

var x []int // nil
// nil is similar, but slightly different from the null in other languages
// nil represents the lack of a value for some types
```

- slices are not comparable
- the only thing you can compare a slice with is nil
  - (my_slice==nil)

>> The reflect package contains a function called DeepEqual that can compare almost anything, including slices. It's primarily intented for testing

Bult in functions in Go to work with slices:

```go
var s = []int{1, 3, 5}
var x = 3
var y = [3]int{7, 9, 11}

fmt.Println(len(s))
s = append(s, x) // Takes at least two parameters, a slice of any type and a value of that type
y = append(s, y...) // ... is a variadic input
```

Go is a call by value language. Every time you pass a parameter to a function, Go makes a copy of the value that is passed in. Padding a slice to the append function actually passes a copy of the slice to the function. The function adds the values to the copy of the slice and returns the copy. You them assign the returned slice back to the variable in the calling function

#### Capacity



