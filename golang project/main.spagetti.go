package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("numbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
