const timeOutMessage = 2000
const myForm = document.querySelector("#form")


myForm.onsubmit = function(e) {
  document.cookie = "token= ;path=/";
  window.location.replace("http://localhost:8000/auth");
  return false
}
