package main

import (
	"flag"
	"fmt"
)

func main() {
	limit := flag.Int("limit", 5, "Number of issue to fetch") // this will give a refrence to the int which will later hold the parsed value

	flag.Parse() //updated the parsed value at *int
	fmt.Printf("Fetchng %d good first issues... \n", *limit)
}
