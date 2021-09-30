package main
import "fmt"

func main() {  
    // declaring array
    a := [5] string {"one", "two", "three", "four", "five"}
    fmt.Println("Array after creation:",a)

    // creating a slice named b
    // first index is inclusive, last index is exclusive
    var b [] string = a[1:4]
    fmt.Println("Slice after creation:",b)

    // changing the slice data
    b[0]="changed"
    fmt.Println("Slice after modifying:",b)
    fmt.Println("Array after slice modification:",a)

    // The length of a slice is the number of elements it contains.
    // The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. 
    fmt.Printf("ARRAY len=%d cap=%d %v\n", len(a), cap(a), a)
    fmt.Printf("SLICE len=%d cap=%d %v\n", len(b), cap(b), b)
}
