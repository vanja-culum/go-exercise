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

	stack := ds.Stack[int]{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.String())

	v, err := stack.Pop()

	if err != nil {
		fmt.Println("err pop", err)
	}

	fmt.Println("pop", v)

	v2, err := stack.Peek()

	if err != nil {
		fmt.Println("peek err", err)
	}

	fmt.Println("peek", v2)

	fmt.Println(stack.String())





    // http.HandleFunc("/hello", helloHandler)
    // http.HandleFunc("/headers", headersHandler)

    // http.ListenAndServe(":8081", nil)

}

