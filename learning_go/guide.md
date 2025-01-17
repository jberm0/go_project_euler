# packages
every go programme is made up of packages, and they start with package `main`

the package name is the last element of the import path, so the package `main` comprises files that start with `package main`

# imports and exports

`import ("fmt", "math")` to import packages

a name is imported if it starts with a capital - e.g. math.pi is not available but math.Pi is

# functions

types are declared after the variable name, and can be grouped. in the below example we could have `x, y int`

if a function returns multiple results, we can return different types with (int, string), for example.

e.g.
func add(x int, y int) int {
	return x + y
}

a function can return named objects declared within it, with a 'naked' return, e.g.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

# variables

`var` defines a list of variables, the type is last. variables can be defined at the function or package level

function level: `var i int` within a function
package level: `var c bool` within the package

if the variable is initialised with a value, the type can be omitted as it is inferred from the initial value

if the variable is not initialised with a value, the variable is given the 0 value - 0, false, "", etc.

inside a function, `:=` can be used for variable assignment in place of `var`

outside a function, every statement must begin with a keyword (`var`, `func`, etc.) so `:=` is not available

# type conversion

the expression T(v) converts the value v to the type T

e.g. float64(3)

conversion in Go must be explicit

# constants

constants are declared with the `const` keyword, and `:=` syntax is not available

# for loop

the foor loop has three components, separated by semicolons
- init statement: executed before the iteration (optional)
- condition expression: evaluated before every iteration
- post statement: executed at the end of every iteration (optional)

the loop stops iterating when the condition expression is false

e.g.
for i := 0; i < 10; i++ {
		sum += i
	}

if no init or post statement is given, the semicolons can be dropped, e.g.
for sum < 1000 {
		sum += sum
	}

# if statement

if {condition} {
    execute ...
}

the if statement can start with a statement to exectute (init statement), any variables declared in this statement are only available in the if condition or an else block

if {condition} {

} else {

}

# switch

a shorter way to write multiple if-elses, it runs the first matching case

switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

a switch with no condition is the same as `switch true`

# defer

a `defer` statement in a function defers the execution of that function until the surrounding function returns

deferred functions are pushed onto a stack, which are executed in last-in-first-out order

# pointers

pointers hold the memory address of a value

the type *T is a pointer to a value of type T

the & operator generates a pointer to its operand - e.g. `p = &i` creates a pointer `p` to the variable `i`

the * operator retrieves or sets the pointer's underlying value - e.g. `fmt.Println(*p)` prints the value of p, `*p = 21` sets the underlying value of `p` to 21 through `i`

# structs

a struct is a collection of fields, whose values can be accessed using `.`

e.g. 

type Vertex struct {
	X, Y int
}

v := Vertex{1, 2}

v.X

structs and struct values can be accessed using pointers

only a subset of values can be given using `name: ` syntax, e.g. v2 = Vertex{X: 1}, the other values are implicit

# arrays

n[T] is an array of n values of type T

an array's length is part of it's type, so arrays cannot be resized

array values are 0-indexed and are retrieved like a[0]