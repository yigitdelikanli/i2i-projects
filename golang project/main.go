package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func open_file(name string) io.Reader {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
func main() {
	name := open_file("numbers.txt")
	scanner := bufio.NewScanner(name)
	for scanner.Scan() {
		strVar := scanner.Text()
		number, _ := strconv.Atoi(strVar)
		var counter = 0
		for i := 2; i < number; i++ {
			if number%i == 0 {
				counter++
				break
			}
		}
		if counter == 0 && number != 1 {
			fmt.Println(number, "is a Prime number")
		} else {
			fmt.Println(number, "is not a prime number")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
