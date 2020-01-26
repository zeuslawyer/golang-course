#README

### CLI Commands

![CLI Commands](/img/CLIcommands.png 'CLI Commands in Go')

### Package types

![Package Types](/img/package-types.png)

- executable packages are always in `package main` and
- files with `package main` _must_ have a `func main(){}` declared

### Lists

Go has two kinds of lists: Arrays (fixed length) and Slice (dynamic length). All elements inside these lists must have an identical type, and so these lists are typed.

#### Append

`cards = append(cards, newCard("5 of Diamonds"))`

#### Iterate

```go
for ind, val := range cards {
  fmt.println(ind, card)
}

```


### testing

- all test files must end with `_test.go`
- the command to run tests is `go test`
- naming:  a func with name `someFunc()` should be tested with `TestSomeFunc()`



### References:

[go syntax by example](https://gobyexample.com/http-clients)
