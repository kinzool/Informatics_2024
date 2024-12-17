package lab9

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (l *ToDoList) AddTask(description string, name string) error {
	task := ToDoTask{Description: description, Status: "не выполнена"}
	for _, el := range l.Tasks {
		if el == task {
			return fmt.Errorf("добавление задачи: задача '%s' уже существует", description)
		}
	}
	l.Tasks = append(l.Tasks, task)
	WriteDataToFile(name, l)
	return nil
}

func (l *ToDoList) DeleteTask(index int, name string) (string, error) {
	if index >= len(l.Tasks) || index < 0 {
		return "", fmt.Errorf("удаление задачи: индекс %d вне диапазона", index)
	}
	output := l.Tasks[index].Description
	l.Tasks = append(l.Tasks[:index], l.Tasks[index+1:]...)
	WriteDataToFile(name, l)
	return output, nil
}

func (l *ToDoList) MarkAsComplete(description string, name string) (string, error) {
	for i := range l.Tasks {
		if strings.Contains(l.Tasks[i].Description, description) {
			l.Tasks[i].Status = "выполнена"
			return l.Tasks[i].Description, nil
		}
	}
	WriteDataToFile(name, l)
	return "", fmt.Errorf("изменение статуса задачи: задачи, соответсвующей описанию, не существует")
}

func (l *ToDoList) SearchTask(keyword string) (string, error) {
	for i := range l.Tasks {
		if strings.Contains(l.Tasks[i].Description, keyword) {
			taskStr, err := json.Marshal(l.Tasks[i])
			if err != nil {
				return "", fmt.Errorf("поиск задачи: %w", err)
			}
			return string(taskStr), nil
		}
	}
	return "", fmt.Errorf("поиск задачи: данная задача не найдена")
}

func (l *ToDoList) PrintToDoList() {
	if len(l.Tasks) == 0 {
		fmt.Println("Нет добавленных задач")
	}
	for i, task := range l.Tasks {
		fmt.Printf("Номер задачи: %d, Описание задачи: %s, Статус задачи: %s\n", i+1, task.Description, task.Status)
	}
}
