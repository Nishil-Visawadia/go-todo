const apiUrl = 'http://localhost:8080/todos';

document.addEventListener('DOMContentLoaded', fetchTodos);

async function fetchTodos() {
    const response = await fetch(apiUrl);
    const todos = await response.json();

    const todoList = document.getElementById('todo-list');
    todoList.innerHTML = ''; // Clear current list

    todos.forEach(todo => {
        const todoItem = document.createElement('li');
        todoItem.innerHTML = `
            ${todo.task} - <span>${todo.status}</span>
            <button class="complete" onclick="updateStatus(${todo.id})">Mark as Completed</button>
            <button class="delete" onclick="deleteTodo(${todo.id})">Delete</button>
        `;
        todoList.appendChild(todoItem);
    });
}

async function addTodo() {
    const taskInput = document.getElementById('task');
    const newTask = taskInput.value;

    if (!newTask) return alert("Please enter a task.");

    const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ task: newTask, status: 'pending' })
    });

    const todo = await response.json();
    fetchTodos(); // Re-fetch the todo list
    taskInput.value = ''; // Clear input field
}

async function updateStatus(id) {
    const response = await fetch(`${apiUrl}/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ status: 'completed' })
    });

    if (response.ok) fetchTodos(); // Re-fetch the todo list
}

async function deleteTodo(id) {
    const response = await fetch(`${apiUrl}/${id}`, {
        method: 'DELETE'
    });

    if (response.ok) fetchTodos(); // Re-fetch the todo list
}
