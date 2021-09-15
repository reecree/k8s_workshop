package main

import (
	"counterd/src/counterd"
	"fmt"
)

func main() {
	r := counterd.InitializeRoutes()
	if r == nil {
		fmt.Printf("No engine received. Returning\n")
		return
	}
	r.Run(":8080")
}
