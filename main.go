package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Record struct {
	ID   string
	Hash string
	Code string
}

var m map[string]Record

func main() {
	file, err := os.OpenFile("value.csv", os.O_RDONLY, 600)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := csv.NewReader(file)
	m = make(map[string]Record)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)

		m[record[1]] = Record{record[0], record[1], record[2]}
	}
	fmt.Println(m)

	file, err = os.OpenFile("in.csv", os.O_RDONLY, 600)
	if err != nil {
		fmt.Println(err)
		return
	}
	r = csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// keyにヒットしない場合空の構造体が作られる？
		if m[record[0]].ID != "" {
			fmt.Printf("%#v\n", m[record[0]])
		}
	}
}
