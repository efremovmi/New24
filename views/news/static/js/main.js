const timeOutMessage = 2000;
const exitFunction = document.getElementById("exit");
const menuAdminFunction = document.getElementById("menu_admin");
const editNewsFunction = document.getElementById("edit_news");



exitFunction.onclick = function(e) {
  document.cookie = "token= ;path=/";
  window.location.replace("http://localhost:8000/auth");
  return false;
}

editNewsFunction.onclick = function(e){
  window.location.replace("http://localhost:8002/news/get-moder-menu");
    return false;
}

menuAdminFunction.onclick = function(e) {
  window.location.replace("http://localhost:8001/control-users");
  return false;
}