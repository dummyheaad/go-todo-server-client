# Todo CLI App
A simple app to create a todo list. The input can be taken from args and STDIN. The created list will then be stored as a JSON file.  
## How to install
### Clone the repository
    git clone https://github.com/dummyheaad/go-cli-todo-app
## Navigate into cmd/todo directory
    cd cmd/todo
### Build the executable
    go build -buildvcs=false .
## Functionalities
### List all tasks
    ./todo -list
### List all tasks with verbose
    ./todo -list -v
### List all uncompleted tasks
    ./todo -list -uncompleted
### List all uncompleted tasks with verbose
    ./todo -list -v -uncompleted
### Add a new task from args
    ./todo -add "my todo"
### Add a new task using STDIN
    echo "my todo" | ./todo -add
### Add multiple tasks using STDIN
    cat << EOF | ./todo -add
    task 1
    task 2
    task 3
    EOF
### Mark a single task as completed (marked with X prefix)
    ./todo -complete task_number
### Delete a single task
    ./todo -del task_number
## Examples
    ./todo -add "task 1"
    echo "task 2" | ./todo -add
    cat << EOF | ./todo -add
    task 3
    task 4
    EOF
    
    ./todo -list
    Output:
      1: task 1
      2: task 2
      3: task 3
      4: task 4
    
    ./todo -list -v
    Output:
      1) Task Name: task 1
         Created At: xxxx
      2) Task Name: task 2
         Created At: xxxx
      3) Task Name: task 3
         Created At: xxxx
      4) Task Name: task 4
         Created At: xxxx
    
    ./todo -complete 1
    ./todo -delete 4
    ./todo -delete 3
    ./todo -list -uncompleted
    Output:
      1: task 1
## Uses different file name as output file
The default file name for the output is .todo.json  
You can also use different name by declaring an environment variable called TODO_FILENAME

    export TODO_FILENAME=my-todo.json