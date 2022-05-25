# Variables
---
Variables have been ask to be nice and clean, we are able to add shadow variables, but not redeclare in the same statement.

- Always have to be use
- Declare variables

  ```
  var i int = 2
  ```
  or to let the compiler decide for us, as an auto type in C++
  
  ```
  i := 2
  ```
  
- Variables in **lower case** will be expose out side outside of the package.

# Primitive
---
Go is very hesitant about implicit data conversion. So, we must be clear of the data type.

## Integer 
- **int8**, **int16**, **int32**, **int64** from different platform which the number represent the bits.  
- **int** is Platform dependent in Golang.  
- **uint** as unsigned integer.  

## Bit operators
- **&, |, ^, &^** bits operator, let's see the example below.

  ```
    a:= 10 // 1010
    b:= 3 // 0011

    // a & b = 0010
    // a | b = 1011
    // a ^ b = 1001
    // a &^ b = oposite of or = 0100
  ```

- **>>, <<** bit shifting

  ```
    a:=8 // 2^3
    // a << 3 = 0100 shift 3 bit left = 2^3 * 2^3 = 2^6
    // a >> 3 = shift 3 bit right = 2^0
  ```

## Floating point numbers   
  Follows IEEE 754 standard, with 32 bit and 64 bit.
  
  ```
    n := 3.14 (var n float32 = 3.14) // Decimal
    n = 13.7e72 // Exponential
    n = 2.1E14 // 13.7e12
  ```
  
## Complext number
  ```
    var n complext64 = 1 + 2i
  ```
  - use real(v) to get the real part (float)
  - use imag(v) to get the imagine part (float)

## string
  - string can be as array just like C++
  - ***immutable***
  - Adding string just use + operator   
  
  Covert the string to byte collection (slice of byte)
  ```
    s := "We are the best"
    b := []byte(s)
    fmt.Printf("%v, %T\n", b, b)
  ```
  
  We get a byte slice with ascii value or utf value, the type result is []unit8  
  ***Many function in Golang is work as slice of byte (string, files)***
  
## rune
  - Golang doesn't have a char data type. It uses byte and rune to represent character values. 
  - Alias for int32
  - Respresent utf32 charater - check this https://pkg.go.dev/strings#Reader.ReadRune
    - if we are working utf32 we can use the above function to get the information we need.   

  ```  
  r := 'a'
  fmt.Printf("%v, %T\n", r, r)
  
  // 97, int32
  ```


# Constant 
---

# Function
---
- https://medium.com/rungo/the-anatomy-of-functions-in-go-de56c050fe11
- Method (Function) in Golang https://go.dev/tour/methods/1

# Array & Slice
---

# Map & Structs
---
## Map
- **Referrence to same underlying data**, not copy.
- Create via literals or make function.
- Check for presence with "value, ok" form of result.

```
	demoMap := make(map[string]int)
	if _, ok := demoMap["a"]; !ok {
		fmt.Println("unfound key a")
	}
```

## Structs 
- Fields can points to any data type in go.
- Normally created as types, but anonymous structs are allowed.
- Structs are value type, will **copy to new struct**.
- Field can be tagged.
- No inheritance, but use **Composition** to embed.
	- Difference between composition and inheritance in Go, is a struct which inherits from another struct can directly access the methods and fields of the parent struct.

```
type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

func (b blogPost) details() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.author.fullName())
	fmt.Println("Bio: ", b.author.bio)
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}

	fmt.Println(author1)
	fmt.Println(author1.fullName())

	blogPost1 := blogPost{
		"This is fantastic",
		"Read through this more than once, promise me.",
		author1,
	}

	fmt.Println(blogPost1)
	fmt.Println(blogPost1.author.fullName())
	blogPost1.details()
}
```
