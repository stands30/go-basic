package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/todoTask/note"
	"example.com/todoTask/todo"
)

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

	todoRef.Display()
	err = todoRef.Save()
	if err != nil {
		fmt.Println("Saving the todo failed")
		fmt.Print(err)
		return
	}
	fmt.Println("Saving the todo succeded")

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		fmt.Print(err)
		return
	}
	fmt.Println("Saving the note succeded")
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
