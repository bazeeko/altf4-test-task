function createTable(tableId) {
    let table = document.getElementById(tableId);

    for (let i = 0; i < 15; i++) {
        let row = table.insertRow();
        row.insertCell(0).innerHTML = i+1;
        row.insertCell(1);
        row.insertCell(2);
    }

    let row = table.insertRow();
    row.insertCell(0);
    row.insertCell(1).innerHTML = "Total";
    row.insertCell(2);
}

function sum(array) {
    let result = 0
    for (let i = 0; i < array.length; i++) {
        result += parseFloat(array[i][1])
    }
    return result
}

function update(book) {
    let bidTable = document.getElementById("bid-table");
    let askTable = document.getElementById("ask-table");

    for (let i = 0; i < book.bids.length; i++) {
        for (let j = 0; j < 2; j++) {
            bidTable.rows[i+1].cells[j+1].innerHTML = Math.round(book.bids[i][j] * 100000) / 100000;
            askTable.rows[i+1].cells[j+1].innerHTML = Math.round(book.asks[i][j] * 100000) / 100000;
        }
    }

    bidTable.rows[16].cells[2].innerHTML = Math.round(sum(book.bids) * 100000) / 100000;
    askTable.rows[16].cells[2].innerHTML = Math.round(sum(book.asks) * 100000) / 100000;
}

let book = {}

const websocket = new WebSocket("ws://localhost:8080/ws")

websocket.onopen = function (event) {
    createTable("bid-table");
    createTable("ask-table");
    console.log("Successfully connected to websocket endpoing");
}

websocket.onerror = function (err) {
    console.log(err);
}

websocket.onmessage = function (event) {
    book = JSON.parse(event.data);
    update(book);
}