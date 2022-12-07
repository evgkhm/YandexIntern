package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Flight struct {
	Day    int
	Hour   int
	Minute int
	Id     int64
	Status string
}

func main() {
	var logCount int
	var err error
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Println("Cannot open file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = -1
	rows, err := reader.ReadAll()
	if err != nil {
		log.Println("Cannot read file:", err)
	}

	var rawData [][]string
	logCount, _ = strconv.Atoi(rows[0][0]) //количество записей в логе
	for i := 1; i <= logCount; i++ {
		//rawData[i-1] = rows[i]
		rawData = append(rawData, rows[i])
	}

	flights := makeFlights(rawData, logCount) //строки с полетами
	unicID := getUnicID(flights, logCount)    //слайс с уникальными ID полетами
	res := calc(flights, unicID, logCount)    //само вычисление
	fmt.Println(res)

	newFile, err := os.Create("output.txt")
	if err != nil {
		log.Println("Cannot create TXT file:", err)
	}
	defer file.Close()
	writer := csv.NewWriter(newFile)
	data := [][]string{
		{res},
	}
	err = writer.WriteAll(data)
	if err != nil {
		log.Println("Cannot write to TXT file:", err)
	}
}

func calc(flights []Flight, unicID []int64, logCount int) string {
	res := 0
	var result string
	for _, val := range unicID {
		for i := 0; i < logCount; i++ {
			if flights[i].Id == val && flights[i].Status == "A" { //начало
				day := flights[i].Day
				hour := flights[i].Hour
				minute := flights[i].Minute
				res -= (day * 24 * 60) + (hour * 60) + minute

				flights = append(flights[:i], flights[i+1:]...) //delete string from slice
				logCount--
				for j := 0; j < logCount; j++ {
					if (flights[j].Status == "S" || flights[j].Status == "C") && flights[j].Id == val &&
						flights[j].Day >= day {
						res += (flights[j].Day * 24 * 60) + (flights[j].Hour * 60) + flights[j].Minute

						flights = append(flights[:j], flights[j+1:]...) //delete string from slice
						logCount--
						break
					}
				}
			} else if flights[i].Status == "B" { //данная строка не нужна
				flights = append(flights[:i], flights[i+1:]...) //delete string from slice
				logCount--
			}
		}
		result += strconv.Itoa(res) + " "
		res = 0
	}
	return result
}

func getUnicID(flights []Flight, logCount int) []int64 {
	m := make(map[int64]int64)
	var idCount []int64
	for i := 0; i < logCount; i++ {
		key := flights[i].Id

		_, ok := m[key]
		if !ok {
			m[key] = flights[i].Id
			idCount = append(idCount, key)
		}
	}
	return idCount
}

func makeFlights(data [][]string, logCount int) []Flight {
	var flights []Flight
	for i := 0; i < logCount; i++ {
		stroka := strings.Join(data[i], " ")
		str := strings.Split(stroka, " ")

		day, hour, minute, id, status := str[0], str[1], str[2], str[3], str[4]
		dayInt, _ := strconv.Atoi(day)
		hourInt, _ := strconv.Atoi(hour)
		minuteInt, _ := strconv.Atoi(minute)
		idInt, _ := strconv.ParseInt(id, 10, 64)
		fly := Flight{
			Day:    dayInt,
			Hour:   hourInt,
			Minute: minuteInt,
			Id:     idInt,
			Status: status,
		}
		flights = append(flights, fly)
	}
	return flights
}
