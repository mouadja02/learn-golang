# Basic variables
## Declaring a variable
```go
var name type = value
```

or 

```go
var name type // value is initialized with the zero value of the type (false for bool, 0 for int, "" for string, 0.0 for float64)
name = value
```

or

### GOATed variable declaration
```go
name := value // type is inferred from the value
```
---

### Basic Types

- bool: a boolean value, can be true or false
- string: a sequence of characters
- int: an integer value
- float64: a floating-point value
- byte: 8-bit of data


### Type Sizes
Integers, uints, floats, and complex numbers all have type sizes.

#### Signed integers (no decimal)
```go
int  int8  int16  int32  int64
```

#### Unsigned integers (non-negative numbers/no decimal)
```go
uint uint8 uint16 uint32 uint64 uintptr
```

#### Signed decimal numbers
```go
float32 float64
```

#### Complex numbers (a complex number has a real and imaginary part)
```go
complex64 complex128
```

### What's the Deal With the Sizes?
The size (8, 16, 32, 64, 128, etc.) represents how many bits in memory will be used to store the variable. The "default" int and uint types refer to their respective 32 or 64-bit sizes depending on the environment of the user.

The "standard" sizes that should be used unless you have a specific performance need (e.g. using less memory) are:

```go
int
uint
float64
complex128
```

### Converting Between Types
Some types can be easily converted like this:

```go
temperatureFloat := 88.26
temperatureInt := int64(temperatureFloat)
```

Casting a float to an integer in this way truncates the floating point portion.