const timeOutMessage = 2000
const exitFunction = document.querySelector("#exit")
const editNews = document.querySelector("#edit_news")


exitFunction.onclick = function(e) {
  document.cookie = "token= ;path=/";
  window.location.replace("http://localhost:8000/auth");
  return false
}