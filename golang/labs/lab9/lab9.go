package lab9

import "fmt"

const filePath = "../golang/labs/lab9/"

var filename string

type ToDoList struct {
	Tasks []ToDoTask `json:"tasks"`
}

type ToDoTask struct {
	Description string `json:"Description"`
	Status      string `json:"Status"`
}

func RunLab9() {
	fmt.Println("Введите имя файла, в который вы хотите записать данные:")
	fmt.Scan(&filename)

	name, err := CreateFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	TaskList := ToDoList{}

	err = ReadDataFromFile(name, &TaskList)
	if err != nil {
		fmt.Println(err)
	}

	err = TaskList.AddTask("Сделать 9 лабу", name)
	if err != nil {
		fmt.Println(err)
	}

	err = TaskList.AddTask("Проснуться к первой паре", name)
	if err != nil {
		fmt.Println(err)
	}

	err = TaskList.DeleteTask(1, name)
	if err != nil {
		fmt.Println(err)
	}

	taskName, err := TaskList.SearchTask("Сделать")
	if err == nil {
		fmt.Println(taskName)
	}
	if err != nil {
		fmt.Println(err)
	}

	TaskList.PrintToDoList()
	err = WriteDataToFile(name, &TaskList)
	if err != nil {
		fmt.Println(err)
		return
	}
}
