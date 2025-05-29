// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"github.com/vanja-culum/go-exercise/ds"
)



func main() {
    lst := ds.List[int]{}
    lst.Append(1)
    lst.Append(2)
	lst.Append(3)
	lst.Append(4)
	lst.Append(5)
	lst.Append(6)
	
	lst.Reverse()

	fmt.Println("reversed",lst.String())

	lst.InsertAt(6, 500)

	fmt.Println(lst.String())

    // http.HandleFunc("/hello", helloHandler)
    // http.HandleFunc("/headers", headersHandler)

    // http.ListenAndServe(":8081", nil)

}

