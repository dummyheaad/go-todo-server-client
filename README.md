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

