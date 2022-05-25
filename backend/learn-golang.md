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
    <br/>
	
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
	<br/>
	
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

# Condition Statement
---
- If Statement

	```
	// num:=9 is an inialization.
	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
	```
- Switch Statement
	Input the tag of switch for checking the condition.
	Tag can be empty input, and also same as If statment with an initializer.
	"break" is already implied.
	
	```
	switch i := 9; i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("unmatched")
	}
	```
	
	Use "fallthrough if we want the statement just continue execute the next statment, be aware it is logicless.

# Looping
---
- Also, with **break** and **continue**

	```
	sum := 0
	for i := 1; i < 5; i++ {
	    sum += i
	}
	```
- Another literal format	

	```
	for idx, value := range LIST {
		...
	}
	
	for key, value := range MAP {
		...
	}
	```
	
- With Loop Label using break and continue.

	```
	guestList := []string{"bill", "jill", "joan", "andy", "kelly"}
	arrived := []string{"sally", "jill", "joan", "kelly"}
	CheckList:
		for _, guest := range guestList {
			for _, person := range arrived {
				fmt.Printf("Guest[%s] Person[%s]\n", guest, person)

				if person == guest {
					fmt.Printf("Let %s In\n", person)
					continue CheckList
				}
			}
		}
	```
