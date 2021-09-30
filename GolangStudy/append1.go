package main
import "fmt"

func main() {  


    type Device struct {
	index         string
	Fab           string
	CXLDevAFUPath string
    }
	a := [5] Device {Device{"un","fab1","/tmp"}, Device{"deux","fab2","/etc"}, Device{"trois","fab3","/home/fabrice"}}
	slice_a := a[1:3]

    fmt.Printf("-----------------------------------------------------------------------------------------------------------------------------------\n")
    fmt.Printf("array a len=%d cap=%d %v\n", len(a), cap(a), a)
    fmt.Printf("slice_a len=%d cap=%d %v\n", len(slice_a), cap(slice_a), slice_a)
    fmt.Printf("-----------------------------------------------------------------------------------------------------------------------------------\n")

    slice_a = append(slice_a,Device{"quatre","fab4","/gudul"}) 

    fmt.Printf("array a len=%d cap=%d %v\n", len(a), cap(a), a)
    fmt.Printf("slice_a len=%d cap=%d %v\n", len(slice_a), cap(slice_a), slice_a)
    fmt.Printf("-----------------------------------------------------------------------------------------------------------------------------------\n")

    slice_a = append(slice_a,Device{"cinq","fab5",""}) 

    fmt.Printf("array a len=%d cap=%d %v\n", len(a), cap(a), a)
    fmt.Printf("slice_a len=%d cap=%d %v\n", len(slice_a), cap(slice_a), slice_a)
    fmt.Printf("-----------------------------------------------------------------------------------------------------------------------------------\n")

    slice_a = append(slice_a,Device{"six","fab6","/dev/sixcarte"}) 

    fmt.Printf("array a len=%d cap=%d %v\n", len(a), cap(a), a)
    fmt.Printf("slice_a len=%d cap=%d %v\n", len(slice_a), cap(slice_a), slice_a)
    fmt.Printf("-----------------------------------------------------------------------------------------------------------------------------------\n")
}
