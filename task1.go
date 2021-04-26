package main

/*
	Напишите функцию для подсчета частоты упоминания слов в строке текста
	и возвращения карты со словами и числом, указывающем, сколько раз они
	употребляются. Функция должна конвертировать текст в нижний регистр и
	обрезать знаки препинания. Используйте пакет strings. Функции, которые
	пригодятся для выполнения данного задания: Fields, ToLower и Trim.
*/

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

const NotExistInMap = 0
const RegExpWords = `[A-Za-zА-Яа-я]+`
const text = "Задача организации, в развитие особенности же развитие различных форм задача же развитие"


func main() {
	wordsSlice := regexp.MustCompile(RegExpWords).FindAllStringSubmatch(toLowerText(text), -1)
	newSlice := updateSlice(wordsSlice)


	intervalLength := len(wordsSlice) / 3

	var wordsMap map[string]int


	wg := new(sync.WaitGroup)

	go countWords1(wordsMap,intervalLength, newSlice, wg)
	go countWords2(wordsMap,intervalLength, newSlice, wg)
	go countWords3(wordsMap,intervalLength, newSlice, wg)

	wg.Wait()

	fmt.Println(wordsMap)
}

func countWords1(myMap map[string]int, interval int, slice []string, wg *sync.WaitGroup){
	wg.Add(1)
	defer wg.Done()

	for i := 0; i < interval; i++ {
		myMap[slice[i]]++
	}
}

func countWords2(myMap map[string]int, interval int, slice []string, wg *sync.WaitGroup){
	wg.Add(1)
	defer wg.Done()

	for i := interval; i < 2 * interval; i++ {
		myMap[slice[i]]++
	}
}

func countWords3(myMap map[string]int, interval int, slice []string, wg *sync.WaitGroup){
	wg.Add(1)
	defer wg.Done()

	for i := 2 * interval; i < 3 * interval; i++ {
		myMap[slice[i]]++
	}
}

func toLowerText(text string) string{
	 text = strings.ToLower(text)
	 return text
}

func updateSlice(sl [][]string) []string{
	newSlice := make([]string, 0)

	for i := 0; i < len(sl); i++ {
		for _, wordValue := range sl[i] {
			newSlice = append(newSlice, wordValue)
		}
	}
	return newSlice
}





