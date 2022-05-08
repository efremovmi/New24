const timeOutMessage = 2000
const myForm = document.querySelector("#form")


myForm.onsubmit = function(e) {
  window.location.replace("http://localhost:8002/news");
  return false;
}
