package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	user := "golu"

	city := fetchCity(user)

	fmt.Println("city name is: ", city)
}

func fetchCity(userName string) string {
	result := ""
	records := readCsvFile("/Users/princeraj/Projects/github.com/octate/csv-parse/test.csv")
	//fmt.Println(records)
	for i := 0; i < len(records); i++ {
		record := records[i]
		if len(record) == 2 {
			if record[0] == userName {
				result = record[1]
				break
			}
		}
	}
	if result == "" {
		result = "not_found"
	}
	return result
}

func readCsvFile(xyz string) [][]string {

	fmt.Println(xyz)
	f, err := os.Open(xyz)
	if err != nil {
		log.Fatal("Unable to read input file "+xyz, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+xyz, err)
	}

	return records
}
