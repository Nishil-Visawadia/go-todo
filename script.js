const apiUrl = 'http://localhost:8080/todos';

// Fetch and display all todos when the page is loaded
document.addEventListener('DOMContentLoaded', fetchTodos);

async function fetchTodos() {
    const response = await fetch(apiUrl);
    const todos = await response.json();

    const todoList = document.getElementById('todo-list');
    todoList.innerHTML = ''; // Clear current list

    todos.forEach(todo => {
        const todoItem = document.createElement('li');
        todoItem.classList.add(todo.status === 'completed' ? 'completed' : 'pending');
        todoItem.setAttribute('data-id', todo.id);

        todoItem.innerHTML = `
            <span>${todo.task}</span>
            <span class="status">${todo.status}</span>
            ${todo.status !== 'completed' ?
                `<button class="complete">Complete</button>` : ''}
            <button class="delete">Delete</button>
        `;
        todoList.appendChild(todoItem);
    });
}

// Add a new task to the todo list
async function addTodo() {
    const taskInput = document.getElementById('task');
    const newTask = taskInput.value.trim();

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

// Event delegation for handling task completion or deletion
document.getElementById('todo-list').addEventListener('click', async (event) => {
    const todoItem = event.target.closest('li');
    if (!todoItem) return;

    const todoId = todoItem.getAttribute('data-id');

    if (event.target.classList.contains('complete')) {
        await updateStatus(todoId, todoItem);
    }

    if (event.target.classList.contains('delete')) {
        await deleteTodo(todoId, todoItem);
    }
});

// Update task status to completed
async function updateStatus(id, todoItem) {
    const response = await fetch(`${apiUrl}/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ status: 'completed' })
    });

    if (response.ok) {
        todoItem.classList.add('completed');
        todoItem.querySelector('.status').textContent = 'completed';
        todoItem.querySelector('.complete').remove();  // Remove the "Complete" button
    }
}

// Delete a todo task from the list
async function deleteTodo(id, todoItem) {
    const response = await fetch(`${apiUrl}/${id}`, {
        method: 'DELETE'
    });

    if (response.ok) {
        todoItem.remove(); // Remove task from the DOM directly
    }
}
