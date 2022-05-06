const timeOutMessage = 2000
const message = document.querySelector(".message")
const editFunction = document.getElementById("edit_button")
const exitFunction = document.getElementById("return_to_main_menu")

exitFunction.onclick = function(e) {
  window.location.replace("http://localhost:8002/news");
  return false
}



editFunction.onclick = function(e) {

var role;

if (document.getElementById('Moderator').checked == true){
  role = 1;
}else if (document.getElementById('User').checked == true){
  role = 0;
}

 const user = {
    username: document.getElementById('login').value,
    password: document.getElementById('pass').value,
    role:     role
  };

  const dataJSON = JSON.stringify(user);

  if (document.getElementById('create').checked == true) {
    axios.post(`http://localhost:8001/control-users/add-user`, dataJSON, {timeout: 5000})
    .then(function (response) {
      if (response.statusText == "OK") {
        message.innerHTML = "Пользователь добавлен!"
        message.classList.add("active")
        setTimeout(() => {
          message.classList.remove("active")
        }, timeOutMessage)

        document.getElementById('login').value = "";
        document.getElementById('pass').value = ""; 

        axios.get("http://localhost:8001/control-users/get-all-users", {timeout: 5000})
        .then(function (response) {
          if (response.statusText == "OK") {
            div = document.getElementById('tbody');
            div.innerHTML =  ``;
            for (let i = 0; i < response.data.users.length; i += 1) {
              console.log(response.data.users[i]);

              var role;
              if (response.data.users[i].role == 2){
                role = "Админ";
              }else if (response.data.users[i].role == 1){
                role = "Модератор";
              }else if (response.data.users[i].role == 0){
                role = "Пользователь";
              }

              div.innerHTML += `<tr>
                                  <td class="text-left">`+response.data.users[i].id+`</td>
                                  <td class="text-left">`+response.data.users[i].user_name+`</td>
                                  <td class="text-left">`+role+`</td>
                                </tr>`;
          }
        }

        })
        .catch(function (error) {
          if (typeof error.response === 'undefined'){
            error = error;
          }else{
            error = error.response.data.err;
          }
          message.innerHTML = error;
          message.classList.add("error", "active")
          setTimeout(() => {
           message.classList.remove("error", "active")
          }, timeOutMessage)
        });
      }

    })
    .catch(function (error) {
      if (typeof error.response === 'undefined'){
        error = error;
      }else{
        error = error.response.data.err;
      }
      message.innerHTML = error;
      message.classList.add("error", "active")
      setTimeout(() => {
       message.classList.remove("error", "active")
      }, timeOutMessage)
    });



  }else if (document.getElementById('delete').checked == true) {
    axios.post(`http://localhost:8001/control-users/delete-user`, dataJSON, {timeout: 5000})
    .then(function (response) {
      if (response.statusText == "OK") {
        message.innerHTML = "Пользователь удален!"
        message.classList.add("active")
        setTimeout(() => {
          message.classList.remove("active")
        }, timeOutMessage)

        document.getElementById('login').value = "";
        document.getElementById('pass').value = ""; 

        axios.get("http://localhost:8001/control-users/get-all-users", {timeout: 5000})
        .then(function (response) {
          if (response.statusText == "OK") {
            div = document.getElementById('tbody');
            div.innerHTML =  ``;
            for (let i = 0; i < response.data.users.length; i += 1) {
              console.log(response.data.users[i]);

              var role;
              if (response.data.users[i].role == 2){
                role = "Админ";
              }else if (response.data.users[i].role == 1){
                role = "Модератор";
              }else if (response.data.users[i].role == 0){
                role = "Пользователь";
              }

              div.innerHTML += `<tr>
                                  <td class="text-left">`+response.data.users[i].id+`</td>
                                  <td class="text-left">`+response.data.users[i].user_name+`</td>
                                  <td class="text-left">`+role+`</td>
                                </tr>`;
          }
        }

        })
        .catch(function (error) {
          if (typeof error.response === 'undefined'){
            error = error;
          }else{
            error = error.response.data.err;
          }
          message.innerHTML = error;
          message.classList.add("error", "active")
          setTimeout(() => {
           message.classList.remove("error", "active")
          }, timeOutMessage)
        });
      }

    })
    .catch(function (error) {
      if (typeof error.response === 'undefined'){
        error = error;
      }else{
        error = error.response.data.err;
      }
      message.innerHTML = error;
      message.classList.add("error", "active")
      setTimeout(() => {
       message.classList.remove("error", "active")
      }, timeOutMessage)
    });
  }
 
  else if (document.getElementById('update').checked == true) {
    axios.post(`http://localhost:8001/control-users/update-user-role`, dataJSON, {timeout: 5000})
    .then(function (response) {
      if (response.statusText == "OK") {
        message.innerHTML = "Пользователь удален!"
        message.classList.add("active")
        setTimeout(() => {
          message.classList.remove("active")
        }, timeOutMessage)

        document.getElementById('login').value = "";
        document.getElementById('pass').value = ""; 

        axios.get("http://localhost:8001/control-users/get-all-users", {timeout: 5000})
        .then(function (response) {
          if (response.statusText == "OK") {
            div = document.getElementById('tbody');
            div.innerHTML =  ``;
            for (let i = 0; i < response.data.users.length; i += 1) {
              console.log(response.data.users[i]);

              var role;
              if (response.data.users[i].role == 2){
                role = "Админ";
              }else if (response.data.users[i].role == 1){
                role = "Модератор";
              }else if (response.data.users[i].role == 0){
                role = "Пользователь";
              }

              div.innerHTML += `<tr>
                                  <td class="text-left">`+response.data.users[i].id+`</td>
                                  <td class="text-left">`+response.data.users[i].user_name+`</td>
                                  <td class="text-left">`+role+`</td>
                                </tr>`;
          }
        }

        })
        .catch(function (error) {
          if (typeof error.response === 'undefined'){
            error = error;
          }else{
            error = error.response.data.err;
          }
          message.innerHTML = error;
          message.classList.add("error", "active")
          setTimeout(() => {
           message.classList.remove("error", "active")
          }, timeOutMessage)
        });
      }

    })
    .catch(function (error) {
      if (typeof error.response === 'undefined'){
        error = error;
      }else{
        error = error.response.data.err;
      }
      message.innerHTML = error;
      message.classList.add("error", "active")
      setTimeout(() => {
       message.classList.remove("error", "active")
      }, timeOutMessage)
    });
  }
  return false
}