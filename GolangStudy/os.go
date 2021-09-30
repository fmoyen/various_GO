package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Stat("/tmp/ostest")
	fmt.Println("err=",err)
	if os.IsNotExist(err) {
	    fmt.Println("IsNotExist if test")
	} else {
	    fmt.Println("IsNotExist else test")
	}

	_, err = os.Stat("/tmp/ostest")
	fmt.Println("err=",err)
	if os.IsExist(err) {
	    fmt.Println("IsExist if test")
	} else {
	    fmt.Println("IsExist else test")
	}

}

/* 
Not obvious behavior for these 2 functions:

[fabrice@FabT480 GolangStudy] touch /tmp/ostest ; go run os.go 
err= <nil>
IsNotExist else test
err= <nil>
IsExist else test
[fabrice@FabT480 GolangStudy] rm /tmp/ostest ; go run os.go 
err= stat /tmp/ostest: no such file or directory
IsNotExist if test
err= stat /tmp/ostest: no such file or directory
IsExist else test

os.Stat reports an error if the file does not exist but does NOT report an error if the file exists (err=nil), so: 

If file exists, Stat does NOT return any error so:
  -> IsNotExist(err) is wrong as err is NOT an error and so does NOT report that the file exists
  -> IsExist(err) is wrong as err is NOT an error and so does NOT report that the file exists (IsExist analyses an error, it does not test a file existence).
       IsExist should be useful and should be true for exemple if you try to create a new file and the file already exists. Then you have a real error because the file (already) exists

If file does not exist, Stat returns an error reporting that the file does not exists so:
  -> IsNotExist(err) is true as err is an error reporting that the file does not exist
  -> IsExist(err) is wrong as err is NOT an error reporting that the file exists.

*/
