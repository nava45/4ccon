package main

import (
	"fmt"
	"log"
	_ "math/rand"
	"time"
)

func g(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go g(i)
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	// Some blocking operation happens here
	var input string
	fmt.Scanln(&input)

	elapsed = time.Since(start)
	log.Printf("GOR Binomial took %s", elapsed)
}
