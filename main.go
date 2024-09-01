package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	todoList := loadCSVData()
	var choice int32

	color.Set(color.FgHiGreen)
	fmt.Print("-\n-------------- To-do List -----------------\n\n")
	displayTodoList(todoList)
	color.Unset()

	color.Set(color.FgHiCyan)
	fmt.Print("-\n-------------- Menu -----------------\n\n")
	fmt.Print("\n 1. Create a new Todo")
	fmt.Print("\n 2. Update an existing Todo")
	fmt.Print("\n 3. Mark a todo as complete")
	fmt.Println()
	color.Unset()

	fmt.Print("\n Choose an option:")
	fmt.Scan(&choice)

	switch {

	case choice == 1:
		createTodo(todoList)

	case choice == 2:
		updateTodo(todoList)

	case choice == 3:
		markTodo(todoList)
	}
}

func loadCSVData() [][]string {

	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()

	fmt.Printf("Test")
	if err != nil {
		panic(err)
	}

	return records
}

func displayTodoList(todolist [][]string) {
	for k, innerSlice := range todolist {
		fmt.Printf("[%d]  ",
			k+1)
		for _, data := range innerSlice {
			fmt.Printf("%s  \n",
				data)
		}
	}
}

func createTodo(todolist [][]string) {

	temp := make([]string, 0)
	fmt.Printf("\n Enter your todo: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	temp = append(temp, strings.TrimSpace(input))

	todolist = append(todolist, temp)

	isWriteSucessfull := writeToCsv(todolist)

	if isWriteSucessfull {
		fmt.Print("Sucessfully saved data to csv!")
	}

}

func updateTodo(todolist [][]string) {

}

func markTodo(todolist [][]string) {

}

func writeToCsv(todolist [][]string) bool {

	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(todolist)
	if err != nil {
		panic(err)
	}

	return true
}
