const form = document.querySelector(".new-contact");
const table = document.querySelector("table");
const btn = document.getElementById("submit-btn");
let fNameValue = document.getElementById("fName");
let lNameValue = document.getElementById("lName");
let phoneNrValue = document.getElementById("phone-nr");
let emailValue = document.getElementById("email");
let otherBtn = document.querySelectorAll('.delete-btn');

async function doRequest() {
  let url = "http://localhost:8080/";
  let res = await fetch(url);

  if (res.ok) {
    let text = await res.text();

    return text;
  } else {
    return `HTTP error: ${res.status}`;
  }
}

doRequest().then((response) => {
  var obj = JSON.parse(response);
  obj.forEach((contactObj) => {
    let id = contactObj["Id"];
    let fName = contactObj["FirstName"];
    let lName = contactObj["LastName"];
    let phoneNr = contactObj["PhoneNumber"];
    let email = contactObj["Email"];
    table.innerHTML += `<tr>
            <td>${fName}</td>
            <td>${lName}</td>
            <td>${phoneNr}</td>
            <td>${email}</td>
            <td><button class="delete-btn" id= "${id}"onClick="delete_click(this.id)">delete</button></td>               
        </tr>`;
  });
});

btn.onclick = function () {
  let request = {
    FirstName: fNameValue.value,
    LastName: lNameValue.value,
    PhoneNumber: phoneNrValue.value,
    Email: emailValue.value,
  };
  fetch("http://localhost:8080/contacts", {
    method: "POST",
    body: JSON.stringify(request),
    headers: {
      "Content-type": "application/json; charset=UTF-8",
    },
  });
  location.reload();
};

function delete_click(clicked_id)
{
  idInt = parseInt(clicked_id)
  let request = {
    Id: idInt,
  }
  fetch("http://localhost:8080/deleteContact", {
    method: "DELETE",
    body: JSON.stringify(request),
    headers: {
      "Content-type": "application/json; charset=UTF-8",
    },
  });
  location.reload();
}