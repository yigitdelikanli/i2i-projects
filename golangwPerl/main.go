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
	readFile, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	return readFile
}

func main() {

	resultTrue := " is a prime \n"
	resultFalse := " is not a prime \n"

	writeFile, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer writeFile.Close()

	var finalData string = ""
	readFile := open_file("numbers.txt")
	scanner := bufio.NewScanner(readFile)
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
			concatenated := fmt.Sprintf("%d%s", number, resultTrue)
			finalData = fmt.Sprintf("%s%s", finalData, concatenated)

		} else {
			concatenated := fmt.Sprintf("%d%s", number, resultFalse)
			finalData = fmt.Sprintf("%s%s", finalData, concatenated)
		}

	}
	error := os.WriteFile("data.txt", []byte(finalData), 0644)
	if error != nil {
		log.Fatal(error)
	}

}
