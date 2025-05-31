// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"github.com/vanja-culum/go-exercise/ds"
)



func main() {
	pq := ds.PriorityQueue[int]{}

	pq.Enqueue(1, 1)
	pq.Enqueue(2, 2)
	pq.Enqueue(3, 3)

	pq.Enqueue(1, 1)
	pq.Enqueue(4, 4)
	pq.Enqueue(5, 5)

	fmt.Println(pq.String())



	pq.Dequeue()
	fmt.Println(pq.String())

    // http.HandleFunc("/hello", helloHandler)
    // http.HandleFunc("/headers", headersHandler)

    // http.ListenAndServe(":8081", nil)

}

