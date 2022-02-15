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

- Floating point numbers   
  Follows E 754

