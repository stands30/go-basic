package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/todoTask/note"
	"example.com/todoTask/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// type outputable interface {
// 	Save() error
// 	Display()
// }

type outputable interface {
	saver
	Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	todoRef, err := todo.New(todoText)

	if err != nil {
		fmt.Print(err)
		return
	}

	outputData(todoRef)
	outputData(userNote)

}

// To print anything use interface{}
func printSomething(value interface{}) {
	fmt.Println(value)
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		fmt.Print(err)
		return err
	}
	fmt.Println("Saving the note succeded")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
