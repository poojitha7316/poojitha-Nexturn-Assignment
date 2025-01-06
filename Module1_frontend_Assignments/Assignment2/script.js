const amountInput = document.getElementById("amount-input");
const descriptionInput = document.getElementById("description-input");
const categoryInput = document.getElementById("category-input");
const expenseTable = document.getElementById("expense-table");

let expenses = JSON.parse(localStorage.getItem("expenses")) || [];

function addExpense() {
    const amount = parseFloat(amountInput.value);
    const description = descriptionInput.value.trim();
    const category = categoryInput.value;

    if (!amount || !description) {
        alert("Please fill out all fields!");
        return;
    }

    expenses.push({ amount, description, category });
    renderExpenses();
}

function deleteExpense(index) {
    expenses.splice(index, 1);
    renderExpenses();
}

function renderExpenses() {
    expenseTable.innerHTML = "";

    expenses.forEach((expense, index) => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${expense.amount.toFixed(2)}</td>
            <td>${expense.description}</td>
            <td>${expense.category}</td>
            <td><button class="btn btn-danger" onclick="deleteExpense(${index})">Delete</button></td>
        `;
        expenseTable.appendChild(row);
    });

    updatetable();
    saveToLocalStorage();
}


function updatetable() {
    const categoryTotals = expenses.reduce((acc, expense) => {
        acc[expense.category] = (acc[expense.category] || 0) + expense.amount;
        return acc;
    }, {});

    

//     // Update the summary table
    const summaryTableBody = document.getElementById("summary-table-body");
    summaryTableBody.innerHTML = ""; // Clear existing table rows
    labels.forEach((label, index) => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${label}</td>
            <td>$${data[index].toFixed(2)}</td>
        `;
        summaryTableBody.appendChild(row);
    });
}



function saveToLocalStorage() {
    localStorage.setItem("expenses", JSON.stringify(expenses));
}

document.getElementById("add-expense-button").addEventListener("click", addExpense);

renderExpenses();
