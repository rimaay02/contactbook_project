const form = document.querySelector(".new-contact");
const table = document.querySelector("table");
const btn = document.querySelector("button");

let fNameValue = document.getElementById("fName");
let lNameValue = document.getElementById("lName");
let phoneNrValue = document.getElementById("phone-nr");
let emailValue = document.getElementById("email");

async function doRequest() {
    let url = 'http://localhost:8080/';
    let res = await fetch(url);

    if (res.ok) {

        let text = await res.text();

        return text;
    } else {
        return `HTTP error: ${res.status}`;
    }
}

doRequest().then(response => {
    var obj = JSON.parse(response);
    obj.forEach(contactObj => {
        let id = contactObj['Id'];
        let fName = contactObj['FirstName'];
        let lName = contactObj['LastName'];
        let phoneNr = contactObj['PhoneNumber'];
        let email = contactObj['Email'];
        table.innerHTML += `<tr>
            <td>${id}</td>
            <td>${fName}</td>
            <td>${lName}</td>
            <td>${phoneNr}</td>
            <td>${email}</td>
            <button>edit</button>
            <button>delete</button>               
        </tr>`;
     });
    
})

btn.onclick = function() {
    let request = {
        fName: fNameValue.value,
        lName: lNameValue.value,
        phoneNr: phoneNrValue.value,
        email: emailValue.value
    }
    fetch("http://localhost:8080/books", {
        method: "POST",
        body: JSON.stringify(request),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
    });
    location.reload();

}