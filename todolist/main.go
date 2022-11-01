package main

import "fmt"

type Todo struct {
	title       string
	description string
	date_i      string
}

type TodoList struct {
	todos []Todo
}

func (tlist *TodoList) addTodo(t Todo) {
	tlist.todos = append(tlist.todos, t)

}

func affTodoList(t TodoList) {

	fmt.Println("Listes des Todos :")
	for i := 0; i < len(t.todos); i++ {
		fmt.Println("N ", i, " -> ", t.todos[i].title, " : ", t.todos[i].description, " : ", t.todos[i].date_i)
	}
}

func main() {
	var td TodoList
	td.addTodo(Todo{"devoir", "pour 3eme", "02/11/2022"})
	td.addTodo(Todo{"coding", "golang", "12/11/2022"})
	td.addTodo(Todo{"coran", "Khaf", "13/11/2022"})
	affTodoList(td)
}
