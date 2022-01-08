# go

🐨 Go lang Study Repository.

## [00 - Basic](https://go.dev/tour/basics/1)

A little bit of the basics.!

### What have learned?

- Programs in go are made up of `packages`.
- Module imports are done using the `import` key and can be factored using parentheses `()` to group multiple imports.
- Named exports are done using campitalyzed words 'SameExample',
- Functions take one or more arguments where the type follows the argument name!
- When arguments share the same type, I can omit all and preserve only the last `x int, y int` for `x, y int`.
- Function can return one or several results.
- Variables can be declared using the `var` key eg `var nameVariable type`.
- Variables can be initialized using `var nameVariable = value`. In this case the type can be omitted as the variable will take the type of the initiator.
- Inside a function I can use short variables using `:=`, but outside the functions I must use the `var` key.
- The basic types are `bool`, `string`, `byte`, `rune`, `float32`, `float64`.
- Zero values ​​are, `0 for numbers`, `false for booleans` and `"" for strings`.
- When a variable is declared without specifying the type, the type is defined according to the assigned value.
- Constants are declared using the `const` key and values ​​cannot be assigned using `:=`.
- Go has only one type of loop, the `for`.
- [Reference](https://go.dev/tour/welcome/1)

## [01 - Hello World](https://github.com/koffelab/go/tree/main/hello)

Hello world in GO!

### What have learned?

- Create a module using the `go init` command.
- Declare a package.
- Import standard packages.
- Importing external packages.
- Synchronize dependencies using the `go mod tidy` command.
- [Reference](https://go.dev/doc/tutorial/getting-started) | [Commit](https://github.com/klaby/go/commit/812b96bf8aa4c0d41e88b5f6e05fee54cabf6431)

## [02 - First module](https://github.com/koffelab/go/tree/main/greetings)

Creating a module!

### What have learned?

- Declaring a function, as well as the type of parameter and return.
- I learned that functions with a capital initial can be called by a function that is not in the same package.
- I learned how to declare a variable and assign the value in one line using `:=`.
- Formatting a string using the `Sprintf` function.
- [Reference](https://go.dev/doc/tutorial/create-module) | [Commit](https://github.com/klaby/go/commit/751aab00017996da2366dbe6feb64b5e00dcb28b)

## [03 - Using module](https://github.com/koffelab/go/tree/main/hello-named)

Import a local module!

### What have learned?

- Import a local module.
- Use the `go mod edit -replace module=directory` command to reference a local module.
- [Reference](https://go.dev/doc/tutorial/call-module-code) | [Commit](https://github.com/klaby/go/commit/7c57b0b3a55113df24a4e1f457ef5730826dfdb4)

## [04 - Handling errors](https://github.com/koffelab/go/tree/main/hello-named)

Validations and error handling!

### What have learned?

- How a function can return multiple values.
- Use the `errors` module to throw exceptions.
- Use the `log` module to display errors in the terminal.
- Assign value to a variable already declared using `=`.
- [Reference](https://go.dev/doc/tutorial/handle-errors) | [Commit](https://github.com/klaby/go/commit/694badb1dd2c5f6adc0e0be4b7c1a9a3153eabbe)

## [05 - Random values](https://github.com/koffelab/go/tree/main/greetings)

Get random values based on a condition!

### What have learned?

- The use of the `init` function to execute functions at startup.
- Declare a `slice/array`.
- Get the size of a data using the `len` function.
- Use the `Seed` function of the `rand` module to initialize the seeds on each run.
- Use the `Intn` function of the `rand` module to generate a random number within the specified maximum size.
- [Reference](https://go.dev/doc/tutorial/random-greeting) | [Commit](https://github.com/klaby/go/commit/65fb916c7b962cac13a8e950569bbbb95b7399ab)

## [06 - Maps, Slices and Iteration](https://github.com/koffelab/go/tree/main/greetings)

Iterate a slice and populate a map!

### What have learned?

- Declare a `map` and initialize using the `make` function.
- Use the `for-range` loop to iterate over a `map/slice/array`.
- Use the blank `_` identifier to ignore the value.
- Use a generic type with `interface {}`.
- [Reference](https://go.dev/doc/tutorial/greetings-multiple-people) | [Commit](https://github.com/klaby/go/commit/a4c424c7bce4b37567543f870050763fe39e32fd)

## [06 - Unit test](https://github.com/koffelab/go/tree/main/greetings)

Unitary tests!

### What have learned?

- A test file must be named with the suffix `_test.go`.
- The test function must start with `Test` name.
- The test function must take a single `t *testing.T` argument providing all the necessary functions for the test.
- Use `go test` command to run.
- Use the `testing` module.
- Use `regexp` module to work with regex.
- [Reference](https://go.dev/doc/tutorial/add-a-test) | [Commit](https://github.com/klaby/go/commit/0438fd2ead28334d735c48c4660c899a0ce51c3d)
