package main

import (
	"bufio"
	"flag"
	"fmt"

	"io"
	"os"

	"strings"

	"pragprog.com/rggo/apis/todo"
)

// Default file name
var todoFileName = ".todo.json"

func main() {
	// Check if the user defined the ENV VAR for a custom file name
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// Parsing command line flags
	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	delete := flag.Int("del", 0, "Item to be deleted")
	verbose := flag.Bool("v", false, "Enable verbose output only for list tasks")
	uncompleted := flag.Bool("uncompleted", false, "Display only uncompleted tasks")

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Here's how to add a new functionality into the tool")
		fmt.Fprintln(flag.CommandLine.Output(), "1. Define the method on todo.go")
		fmt.Fprintln(flag.CommandLine.Output(), "2. Define the interface on main.go such as flag and switch case")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Define an items list
	l := &todo.List{}

	// Use the Get method to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the provided flags
	switch {
	case *add:
		// When any arguments (excluding flags) are provided, they will be
		// used as the new task
		tasks, err := getTasks(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, t := range tasks {
			l.Add(t)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *complete > 0:
		// Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *delete > 0:
		// Delete the given item
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *list:
		// List current to do items
		switch {
		case *uncompleted && *verbose:
			l = l.FilterUncompleted()
			l.PrintVerbose()
		case *uncompleted:
			l = l.FilterUncompleted()
			fmt.Print(l)
		case *verbose:
			l.PrintVerbose()
		default:
			fmt.Print(l)
		}

	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

// getTask function decides where to get the description for a new
// task from: arguments or STDIN
func getTasks(r io.Reader, args ...string) ([]string, error) {
	// new task from argument (example: ./todo -add "my new todo")
	if len(args) > 0 {
		return []string{strings.Join(args, " ")}, nil
	}

	// new task from STDIN
	var tasks []string
	s := bufio.NewScanner(r)

	for s.Scan() {
		text := strings.TrimSpace(s.Text())
		if len(text) > 0 {
			tasks = append(tasks, text)
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, fmt.Errorf("task cannot be blank")
	}

	return tasks, nil
}
