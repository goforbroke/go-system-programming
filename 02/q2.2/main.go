package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	f, err := os.Create("result.csv")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(f)

	records := [][]string{
		{"foo", "bar"},
		{"fuga", "hoge"},
	}

	for _, r := range records {
		if err := writer.Write(r); err != nil {
			log.Fatal("Somthing happens.")
		}
	}
	writer.Flush()
}
