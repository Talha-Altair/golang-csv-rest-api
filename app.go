package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

type empData struct {
    id string
    state string
    city string
}

func read() {

    csvFile, err := os.Open("./data.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()
    
    csvLines, err := csv.NewReader(csvFile).ReadAll()

    if err != nil {
        fmt.Println(err)
    }

	data := {}

    for _, line := range csvLines {

        emp := empData{
            id: line[0],
            state: line[1],
			city: line[2],
        }

		data.append(emp)

        fmt.Println(emp.state)
    }

	fmt.Println(data)
}

func main() {
	read()
}