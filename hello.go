package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tandahunter/stringutil"
)

func main() {
	sayHello()
	printFirstArgument()
	printAllArguments()
	printAllArgumentsFast()
}

func sayHello() {
	fmt.Println(stringutil.Reverse("!oG, olleH"))
}

func printFirstArgument() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	} else {
		fmt.Println("No args")
	}
}

func printAllArguments() {
	start := time.Now()
	for _, a := range os.Args[1:] {
		fmt.Printf("%s ", a)
	}

	fmt.Printf("%d nanoseconds", time.Since(start).Nanoseconds())
	fmt.Println()
}

func printAllArgumentsFast() {
	start := time.Now()
	fmt.Printf("%s", strings.Join(os.Args[1:], " "))
	fmt.Printf("%d nanoseconds", time.Since(start).Nanoseconds())
}
