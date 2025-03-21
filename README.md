# Go Todo App Server Client
A repo that contains client, server, and storage for todo app build using Golang  
Server: REST API  
Client: Cobra CLI  
Storage: JSON  

## How to install
### Clone the repository

    git clone https://github.com/dummyheaad/go-todo-server-client.git

### Build the executable (in todo, todoClient, and todoServer)

    go build

### Run the server (in todoServer dir)

    ./todoServer

## Functionalities (make sure the server has been running)
### Adding a new task

    ./todoClient add NEW_TASK

### List all tasks

    ./todoClient list

### List only active (not completed) tasks

    ./todoClient --active

### View detail of a single task

    ./todoClient view TASK_NUMBER

### Mark a task as completed

    ./todoClient complete TASK_NUMBER

### Delete a task

    ./todoClient delete TASK_NUMBER

### Generate markdown documentation

    ./todoClient docs --dir PATH

### Generate bash auto-completion

    script <(./todoClient completion)

### For more details, consult the /todoClient/docs directory

## Examples
### Run the server (todo will be saved at /tmp/testtodoclient01.json)

    ./todoServer -f /tmp/testtodoclient01.json

### Add new task

    ./todoClient add my new task
    ./todoClient add my new task 1

### List task

    ./todoClient list
    Output:
    - 1 my new task
    - 2 my new task 1

### Complete the 1st task

    ./todoClient complete 1
    Output:
    Item number 1 marked as completed.

### List task

    ./todoClient list
    Output:
    x 1 my new task
    - 2 my new task 1

### List all active task (not completed)

    ./todoClient list --active
    Output:
    - 2 my new task 1

### Delete a single task
    
    ./todoClient del 1
    Output:
    Item number 1 deleted.

### List task

    ./todoClient list
    Output:
    - 1 my new task 1