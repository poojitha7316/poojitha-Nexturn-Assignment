// script.js

// DOM Elements
const taskInput = document.getElementById('taskInput');
const addTaskButton = document.getElementById('addTaskButton');
const taskListContainer = document.getElementById('taskListContainer');
const pendingCount = document.getElementById('pendingCount');

// Load tasks from localStorage
let tasks = JSON.parse(localStorage.getItem('tasks')) || [];

// Function to update the task list and save to localStorage
function updateTaskList() {
    // Clear the task list container
    taskListContainer.innerHTML = '';
    
    // Create task elements dynamically
    tasks.forEach((task, index) => {
        const taskDiv = document.createElement('div');
        taskDiv.classList.add('task');
        if (task.completed) taskDiv.classList.add('completed');
        
        taskDiv.innerHTML = `
            <span>${task.name}</span>
            <div>
                <button onclick="toggleCompletion(${index})">✓</button>
                <button onclick="editTask(${index})">✏️</button>
                <button onclick="deleteTask(${index})">🗑️</button>
            </div>
        `;
        
        taskListContainer.appendChild(taskDiv);
    });
    
    // Update the pending tasks count
    const pendingTasks = tasks.filter(task => !task.completed).length;
    pendingCount.textContent = `${pendingTasks} Pending Task(s)`;
    
    // Save tasks to localStorage
    localStorage.setItem('tasks', JSON.stringify(tasks));
}

// Add a new task
function addTask() {
    const taskName = taskInput.value.trim();
    if (taskName) {
        tasks.push({ name: taskName, completed: false });
        taskInput.value = ''; // Clear the input field
        updateTaskList();
    }
}

// Toggle completion status of a task
function toggleCompletion(index) {
    tasks[index].completed = !tasks[index].completed;
    updateTaskList();
}

// Edit an existing task
function editTask(index) {
    const newTaskName = prompt("Edit Task", tasks[index].name);
    if (newTaskName !== null) {
        tasks[index].name = newTaskName.trim();
        updateTaskList();
    }
}

// Delete a task
function deleteTask(index) {
    tasks.splice(index, 1);
    updateTaskList();
}

// Event Listeners
addTaskButton.addEventListener('click', addTask);

// Load tasks initially
updateTaskList();


