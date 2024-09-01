package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {

	todoList := loadCSV()
	var choice int32

	fmt.Print("-\n-------------- To-do List -----------------\n\n")
	displayTodoList(todoList)
	fmt.Print("-\n-------------- Menu -----------------\n\n")
	fmt.Print("\n 1. Create a new Todo")
	fmt.Print("\n 2. Update an existing Todo")
	fmt.Print("\n 3. Mark a todo as complete")
	fmt.Println()
	fmt.Print("\n 3. Mark a todo as complete")

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

func loadCSV() map[int]string {

	todolist := make(map[int]string)

	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	var newSlice []string = records[1:][1:]

	for r := range newSlice {

	}
	if err != nil {
		panic(err)
	}

	for _, records := range records {

		if len(records) < 2 {
			continue
		}

		key, err := strconv.Atoi(records[0])
		if err != nil {
			panic(err)
		}

		value := records[1]

		todolist[key] = value
	}

	return todolist
}

func displayTodoList(todolist map[int](string)) {
	for k, v := range todolist {

		fmt.Printf("%d     %s\n",
			k, v)
	}
}

func createTodo(todolist map[int](string)) {

	fmt.Printf("\n Enter your todo: ")
	todo, err := fmt.Scanf()
	if err != nil {
		panic(err)
	}

	lastKey := len(todolist)
	key := lastKey + 1

	todolist[key] = todo
}

func updateTodo(todolist map[int](string)) {

}

func markTodo(todolist map[int](string)) {

}

func writeToCsv() {

}

func getKey() int {

}
