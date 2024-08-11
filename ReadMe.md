if u have no packages can use

```
go mod init <module_name>
```

### dataType
The unsafe.Sizeof function is very useful for understanding how much memory different variables consume.
The size of int is platform-dependent and can be checked using bits.UintSize.
Floating-point numbers have fixed sizes (float32 and float64), which are consistent across platforms.


those ones that already exists its values are value type and the one  that refrence exists at stack are refrence type

```
 var jp *int = &j 
```
Pointer Type *int:

The asterisk (*) before the int indicates that jp is a pointer variable that points to an integer type

### go-cli

go bug
go build --> make exec file
go clean
go doc --> eg go doc fmt
go env


### gc
go has heap and stack that point to them
and gc works with heap to find  ones that need to be cleaned


Always prefer a rune slice when dealing with strings that might contain multi-byte characters to prevent misinterpretation of characters.
Go standard libraries provide robust support for Unicode, allowing developers to manipulate text in a multilingual context easily.


### python 
uses pass by assignment