> Generic function

```go
func PrintAnything[T any](thing T) {
    fmt.Println(thing)
}

// how to use it
PrintAnything[T any](variable)
```
<br>

> Constraints

Take a slice of some element type E, and returns a single string that joins them all together

```go
func Join[T any](things []T) (result string) {
    for _, v := range things {
        result += v.String()
    }
    return result
}

// how to use it
output := Join([]string{"a", "b", "c"})
```

Instead of accepting literally any T, 
we're really only interested in types that have a String() method.
We can have this by using the Stringer interface

```go
func Join[T Stringer](things []T) (result string) {
    for _, v := range things {
        result += v.String()
    }
    return result
}

// how to use it
output := Join([]string{"a", "b", "c"})
```
<br>

> Comparable constraint

Suppose we want to write an Equal function that takes two parameters of type T, 
and return true if they are equal, or false otherwise.

```go
// This won't work.
func Equal[T any](a, b T) bool {
    return a == b
}

// how to use it
fmt.Println(Equal(1, 1))
// cannot compare a == b (operator == not defined for T)
```

This is the same kind of issue as we had in Join() 
with the String() method, but since we're not calling 
a method now, we can't use a constraint based on a method set. 
Instead, we need to constrain T to only types that work 
with the == or != operators, which are known as comparable types. 
Fortunately there's a straightforward way to specify this: 
use the built-in comparable constraint, instead of any.

```go
// This will work.
func Equal[T comparable](a, b T) bool {
    return a == b
}

// how to use it
fmt.Println(Equal(1, 1))
```

> Constraints package

Just to be difficult, 
suppose we want to do something with values of some 
unspecified type that isn't either comparing them or 
calling methods on them.

Suppose we want to write a Max() function 
for a generic type T that takes a slice of T 
and returns the element with the highest value. 
We might try something like this:

```go
func Max[E any](input []E) (max E) {
    for _, v := range input {
        if v > max {
            max = v
		}
	}
    return max
}

// how to use it
fmt.Println(Max([]int{1, 2, 3}))
// cannot compare v > max (operator > not defined for E)
```

Again, Go can't prove ahead of time that the type T 
will work with the > operator 
(we would say that T is ordered, in that case).

```go
type Ordered interface {
    int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string
}
```

```go
func Max[E Ordered](input []E) (max E) {
    for _, v := range input {
        if v > max {
            max = v
		}
	}
    return max
}

// how to use it
fmt.Println(Max([]int{1, 2, 3}))
// cannot compare v > max (operator > not defined for E)
```

Here, we've specified a union of type literals (like int), separated by | (pipe) characters. 
Any of the types in this set will match the constraint Ordered.

Even this wouldn't be enough for a useful Max function, 
though, because it wouldn't include any types derived 
from int, int8, and so on. 
For example, if we defined something like this:

```go
type MyInt int
```

Then we wouldn't be able to use MyInt with our Ordered constraint, 
because it's not in the list!

Instead, we can use a type approximation, 
written with the ~ (tilde) symbol:

```go
type Ordered interface {
    ~int | ~int8 | ~int16 ... // and so on
}
```

```go
func Max[E Ordered](input []E) (max E) {
    for _, v := range input {
        if v > max {
            max = v
		}
	}
    return max
}

// how to use it
fmt.Println(Max([]int{1, 2, 3}))
// cannot compare v > max (operator > not defined for E)
```

A type approximation such as ~int matches not only int itself, 
as you'd expect, but also any type derived from int (like MyInt). 
And if we approximate all the built-in numeric types in this way (and include ~string, which is also ordered), 
we'll have the constraint we need to be able to use the > operator.

Fortunately for your keyboard, 
this is already defined for us in the constraints package, 
along with many others; 
this one is called constraints.Ordered. 
Here's how we use it:

```go
func Max[E constraints.Ordered](input []E) (max E) {
    for _, v := range input {
        if v > max {
            max = v
		}
	}
    return max
}

// how to use it
fmt.Println(Max([]int{1, 2, 3}))
// cannot compare v > max (operator > not defined for E)
```

> Generic types

We know how to write functions that can take arguments of any type. 
But what if we want to create a type that can contain any type? 
For example, a 'slice of anything' type. 
That turns out to be very easy:

```go
type Bunch[T any] []T
```

We're saying that for any given element type T, 
a Bunch[T] is a slice of values of type T. 
For example, a Bunch[int] is a slice of int. 
We can create values of that type in the normal way:

```go
x := Bunch[int]{1, 2, 3}
```

Just as you'd expect, 
we can write generic functions that take generic types:

```go
func PrintBunch[T any](b Bunch[T]) {
```

Methods, too:

```go
func (b Bunch[T]) Print() {
```

We can also apply constraints to generic types:

```go
type StringableBunch[T Stringer] []T
```