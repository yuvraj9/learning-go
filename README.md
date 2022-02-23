# learning-go
Learning go programming by completing default programs present in go documentation and reading Learning Go book.

## Go documentation
This contains two modules - hello and greetings. Hello imports greetings module to return random greetings according to input. More info - https://golang.org/doc/tutorial/greetings-multiple-people

---
## Learning Go: An idiomatic approach to real world Go programming
---
### Chapter 1 - Go Basics

**Effective Go** - https://go.dev/doc/effective_go

**Code review comments** - https://github.com/golang/go/wiki/CodeReviewComments

**Linting** - golangci-lint - https://golangci-lint.run/usage/quick-start/

---
### Chapter 2 - Primitive types

 - Can be declared using `var` and `:=`.
```
    var x int = 10
    var x, y = 10, "hello"
```
- `:=` can be used only in a function. We can assign a new value to a variable by using `:=` untill there is a new variable on the left hand side.
```
    x := 10
    x, y := 30, 20
```
- By default zero value is assinged to the variable. Also prefer using var form when assigning a type of variable
```
    var x int
    x := byte(20) // This is valid but to idiomatic.
    var x byte = 20 // This is idiomatic
```
- Rarely declare variables outside of functions (i.e. in package block). It makes data flow analysis complicated.
- It is allowed to create unread package level variables.
- There can be unread const as they are only used during compile time and wont be included in binary.
- camel case is preferred while naming variables. Also keep the name short for variables declared inside a function. One letter can be used as well like - f for float, i for int, b for boolean.