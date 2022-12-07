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

type Potions struct {
	PotionsQty   int64      //количество зелий (5 зелий 2 чистых)
	Potion       [][]string //состав изделия (А, А, В...)
	QuestionsQty int64      //кол-во вопросов
	Questions    [][]string
}

func main() {
	potions := Potions{}

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

	var i int64
	potions.PotionsQty, _ = strconv.ParseInt(rows[i][0], 10, 64) //количество зелий
	potions.PotionsQty -= 2                                      //убираем чистые зелья
	i++

	for j := i; i <= potions.PotionsQty; j++ {
		potions.Potion = append(potions.Potion, rows[j])
		i++
	}

	potions.QuestionsQty, _ = strconv.ParseInt(rows[i][0], 10, 64) //количество вопросов
	i++

	for j := i; j < potions.QuestionsQty+potions.PotionsQty+2; j++ {
		potions.Questions = append(potions.Questions, rows[j])
		i++
	}

	m := parsePotions(potions)
	res := calc(m, potions)
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

func calc(potionMap map[int64]string, potions Potions) string {
	res := ""
	var i int64
	for i = 0; i < potions.QuestionsQty; i++ {
		stroka := strings.Join(potions.Questions[i], " ")
		str := strings.Split(stroka, " ")
		aQuantityQuestion, _ := strconv.ParseInt(str[0], 10, 64)
		bQuantityQuestion, _ := strconv.ParseInt(str[1], 10, 64)
		whichPotion, _ := strconv.ParseInt(str[2], 10, 64)

		var aQuantityPotionMap int64 = 0
		var bQuantityPotionMap int64 = 0
		potionString := potionMap[whichPotion]
		for _, runa := range potionString {
			if runa == 'A' {
				aQuantityPotionMap++
			} else if runa == 'B' {
				bQuantityPotionMap++
			}
		}

		/*if aQuantityQuestion == 0 || bQuantityQuestion == 0 {
			res += "0"
			continue
		}*/

		_, ok := potionMap[whichPotion]
		if !ok {
			res += "0"
			continue
		}
		val := potionMap[whichPotion]
		if val == "" {
			res += "0"
			continue
		}

		if aQuantityQuestion > 0 && aQuantityPotionMap == 0 { //костыль
			res += "0"
			continue
		} else if bQuantityQuestion > 0 && bQuantityPotionMap == 0 {
			res += "0"
			continue
		}
		/*if whichPotion <= 0 { //костыль
			res += "0"
			continue
		}*/

		if aQuantityQuestion >= aQuantityPotionMap && bQuantityQuestion >= bQuantityPotionMap {
			res += "1"
		} else {
			res += "0"
		}
	}
	return res
}

func parsePotions(potions Potions) map[int64]string {
	var newPotion int64 = 3
	m := make(map[int64]string)
	var i int64
	for i = 0; i < potions.PotionsQty; i++ {
		stroka := strings.Join(potions.Potion[i], " ")
		str := strings.Split(stroka, " ")

		compositPotionQty, _ := strconv.ParseInt(str[0], 10, 64)
		var j int64
		for j = 1; j <= compositPotionQty; j++ {
			if str[j] == "1" {
				m[newPotion] += "A"
			} else if str[j] == "2" {
				m[newPotion] += "B"
			} else {
				key, _ := strconv.ParseInt(str[j], 10, 64)
				m[newPotion] += m[key]
			}
		}
		newPotion++
	}
	return m
}
