# Pointer

j& => the address of a variable

In Go, (*int) and (*p) are both pointers, but they point to different types of data.

\*int is a pointer to an integer. It means that the variable holds the memory address of an integer value. For example:

```go
var a int = 10
var p *int = &a

```

In this case, p is a pointer to an integer, and it holds the memory address of variable a.

On the other hand, \*p is used to dereference the pointer p, i.e., it gives you the value stored at the memory address held by p.

p is a pointer to an integer, *p will be an integer. If p is a pointer to a custom type,*p will be a variable of that custom type. For example:

```go
type MyType struct {
    x int
}

var p *MyType = &MyType{10}
var value = (*p).x  // value will be 10
```

p is a pointer to an integer, *p will be an integer. If p is a pointer to a custom type,*p will be a variable of that custom type. For example:

```go
var p int = 10
fmt.Println(&p)  // This will print the memory address of p
```

## when do we user *int and*p

(*int) and (*p) are used in different contexts:

\*int: This is used when you want to declare a variable that is a pointer to an integer. This means that the variable will hold the memory address of an integer value. For example:
