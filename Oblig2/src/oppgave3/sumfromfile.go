package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
//Ved bruk av funksjonen readLines kjøres da programmet til å lese linjene i resultat.txt
func sumfromfile()  {
	lines, err := readLines("resultat.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	//Deretter tar programmet og legger dem sammen
	input1 := lines[0]
	input2 := lines[1]
	nummer1,_ := strconv.Atoi(input1)
	nummer2,_ := strconv.Atoi(input2)

	resultat := nummer1 + nummer2
	if resultat <= 0 {
		fmt.Println("Error, du må skrive inn tall som er høyere enn 0")
	}
	f, err := os.OpenFile("resultat.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fprintf(f,"\n%d\n", resultat); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
