package main

import (
	"fmt"
	"os"
)

func main() {
	echo(os.Args)
	fileOpen(os.Args)
	findDups()
}

func echo(args []string) {

	sep := ""

	for index, value := range os.Args {
		value += sep + os.Args[index]
		sep = " "
	}

	fmt.Println(s)
}

func check(s string, err error) {
	if err != nil {
		fmt.Println("%s Error: %s", s, err)
	}
	return
}

func fileOpen(args []string) {
	var err error
	var n int
	var f *os.File

	f, err = os.Create(args[1])
	check("Open", err)
	defer f.Close()

	n, err = f.WriteString("Hello World")
	check("Write", err)

	f.Close()

	f, err = os.Open(args[1])
	check("Open2", err)

	b := make([]byte, 50)

	n, err = f.Read(b)
	check("Read", err)

	fmt.Printf("Success: %s Len: %d \n", string(b), n)
}

func findDups() {
	s := ""
	counts := make(map[string]int)

	for i := 0; i < 10; i++ {
		fmt.Scanf("%s", &s)
		if s == "!" {
			break
		}
		counts[s]++
	}

	for line, n := range counts {
		fmt.Println(n, line)
	}
}
