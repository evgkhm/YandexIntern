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

type User struct {
	FirstName  string
	LastName   string
	Patronymic string
	Day        int
	Month      int
	Year       int
}

func main() {
	var usersNum int
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

	var users []User
	for ind, row := range rows {
		if ind == 0 {
			usersNum, _ = strconv.Atoi(row[0])
			continue
		}

		day, _ := strconv.Atoi(row[3])
		month, _ := strconv.Atoi(row[4])
		year, _ := strconv.Atoi(row[5])
		user := User{
			LastName:   row[0],
			FirstName:  row[1],
			Patronymic: row[2],
			Day:        day,
			Month:      month,
			Year:       year,
		}
		users = append(users, user)
	}

	res := Calc(users, usersNum)

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

func Calc(users []User, num int) string {
	var res string
	for i := 0; i < num; i++ {
		unicChars := CalcUnicChars(users[i].LastName, users[i].FirstName, users[i].Patronymic)
		sumDayMonth := CalcSumDayMonth(users[i].Day, users[i].Month)
		alphabetNum := GetAlphNum(users[i].LastName)
		cipher := unicChars + (sumDayMonth * 64) + (alphabetNum * 256)
		hexCipher := strconv.FormatInt(int64(cipher), 16)

		res += lastThreeDigit(hexCipher) + " "
	}
	res = strings.ToUpper(res)
	fmt.Println(res)
	return res
}

// lastThreeDigit только последние 3 цифры
func lastThreeDigit(cipher string) string {
	if len(cipher) > 3 {
		var res string
		length := len(cipher)
		for i := length - 3; i < length; i++ {
			res += string(cipher[i])
		}
		return res
	}
	return cipher
}

// GetAlphNum в алфавите первой буквы фамилии
func GetAlphNum(name string) int {
	num := name[0] - 64
	return int(num)
}

// CalcSumDayMonth сумма цифр в дне и месяце рождения
func CalcSumDayMonth(day int, month int) int {
	res := (day / 10) + (day % 10)
	res += (month / 10) + (month % 10)
	return res
}

// CalcUnicChars узнаем количество уникальных симолов в ФИО
func CalcUnicChars(firstName string, lastName string, patronymic string) int {
	m := make(map[string]int)
	var count int
	for _, runa := range firstName {
		_, ok := m[string(runa)]
		if !ok {
			m[string(runa)] = count
			count++
		}
	}

	for _, runa := range lastName {
		_, ok := m[string(runa)]
		if !ok {
			m[string(runa)] = count
			count++
		}
	}

	for _, runa := range patronymic {
		_, ok := m[string(runa)]
		if !ok {
			m[string(runa)] = count
			count++
		}
	}

	return count
}
