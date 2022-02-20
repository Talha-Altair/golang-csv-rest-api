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



    for _, line := range csvLines {

        emp := empData{
            id: line[0],
            state: line[1],
			city: line[2],
        }

        fmt.Println(emp.state)
    }
}

func main() {
	read()
}