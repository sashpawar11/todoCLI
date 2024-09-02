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

	var choice int32
	exit := false

	for !exit {
		todoList := loadCSVData()
		color.Set(color.FgHiGreen)
		fmt.Print("-\n-------------- To-do List -----------------\n\n")
		displayTodoList(todoList)
		color.Unset()

		color.Set(color.FgHiCyan)
		fmt.Print("-\n-------------- Menu -----------------\n\n")
		fmt.Print("\n 1. Create a new Todo")
		fmt.Print("\n 2. Update an existing Todo")
		fmt.Print("\n 3. Mark a todo as complete")
		fmt.Print("\n 4. Exit")
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

		case choice == 4:
			exit = true
			os.Exit(1)
		}
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

	fmt.Printf("\n Enter the number of the todo you want to update: ")
	var idx int32
	fmt.Scan(&idx)

	fmt.Printf("\n Enter updated todo: ")
	updatedString, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	todolist[idx-1][0] = strings.TrimSpace(updatedString)
	isWriteSucessfull := writeToCsv(todolist)
	if isWriteSucessfull {
		fmt.Print("Sucessfully updated data to csv!")
	}
}

func markTodo(todolist [][]string) {
	fmt.Printf("\n Enter the number of the todo you want to mark as completed : ")
	var idx int32
	fmt.Scan(&idx)

	idx = idx - 1
	fmt.Print(idx)
	todolist = append(todolist[:idx], todolist[idx+1:]...)

	isWriteSucessfull := writeToCsv(todolist)
	if isWriteSucessfull {
		fmt.Print("Todo marked as complete!")
	}

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
