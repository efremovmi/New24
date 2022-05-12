const timeOutMessage = 2000
const myForm = document.querySelector("#form")


myForm.onsubmit = function(e) {
  window.location.replace("http://localhost/news");
  return false;
}
