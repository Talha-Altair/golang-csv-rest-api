package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

type cityData struct {
    id string
    state string
    city string
}

func read() []cityData {

    csvFile, err := os.Open("./data.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()
    
    csvLines, err := csv.NewReader(csvFile).ReadAll()

    if err != nil {
        fmt.Println(err)
    }

	data := []cityData{}

    for _, line := range csvLines {

        city := cityData{
            id: line[0],
            state: line[1],
			city: line[2],
        }

		if city.id == "id" {
			continue
		}

		data = append(data, city)
    }

	return data
}