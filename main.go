// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"github.com/vanja-culum/go-exercise/ds"
)



func main() {
	tree := ds.BST[int]{}

	rootNode := tree.Insert(5)
	rootRight := tree.Insert(10)
	rootLeft := tree.Insert(2)
	fmt.Println(tree.StringRoot())
	fmt.Println("n1", rootNode)
	fmt.Println("n2", rootRight)
	fmt.Println("n3", rootLeft)
	fmt.Println(tree.String())
    // http.HandleFunc("/hello", helloHandler)
    // http.HandleFunc("/headers", headersHandler)

    // http.ListenAndServe(":8081", nil)

}

