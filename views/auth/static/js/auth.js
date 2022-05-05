const timeOutMessage = 2000
const message = document.querySelector(".message")
const myForm = document.querySelector("#form")


myForm.onsubmit = function(e) {
 const user = {
    username: document.getElementById('email').value,
    password: document.getElementById('pass').value
  };

  const dataJSON = JSON.stringify(user);

  if (document.getElementById('signup').checked == true) {
    if (document.getElementById('pass').value != document.getElementById('repass').value) {
      message.innerHTML = "Пароли не совпадают"
      message.classList.add("error", "active")
      setTimeout(() => {
          message.classList.remove("error", "active")
      },  timeOutMessage)
      return false
    }

    axios.post(`http://localhost:8000/auth/sign-up`, dataJSON, {timeout: 5000})
    .then(function (response) {
      if (response.statusText == "OK") {
        message.innerHTML = "Вы успешно зарегестрировались!"
        message.classList.add("active")
        setTimeout(() => {
          message.classList.remove("active")
        }, timeOutMessage)

        document.getElementById('email').value = "";
        document.getElementById('pass').value = "";
        document.getElementById('repass').value = "";
        document.getElementById('signup').checked = false;
        document.getElementById('signin').checked = true;
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
  }else{
    axios.post(`http://localhost:8000/auth/sign-in`, dataJSON, {timeout: 5000})
    .then(function (response) {
      console.log(response.data.token)
      if (response.statusText == "OK") {
        document.cookie = encodeURIComponent("token") + '=' + encodeURIComponent(response.data.token)+ '; samesite=strict; secure';
        message.innerHTML = "Вы успешно вошли!"
        message.classList.add("active")
        setTimeout(() => {
          message.classList.remove("active")
        }, timeOutMessage)

        document.getElementById('email').value = "";
        document.getElementById('pass').value = "";
        document.getElementById('repass').value = "";
        document.getElementById('signup').checked = false;
        document.getElementById('signin').checked = true;
      }
      window.location.replace("http://localhost:8002/news");
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