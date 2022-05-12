const timeOutMessage = 2000;
const exitFunction = document.getElementById("exit");
const menuAdminFunction = document.getElementById("menu_admin");
const editNewsFunction = document.getElementById("edit_news");



exitFunction.onclick = function(e) {
  document.cookie = "token= ;path=/";
  window.location.replace("http://localhost/auth");
  return false;
}

editNewsFunction.onclick = function(e){
  window.location.replace("http://localhost/news/get-moder-menu");
    return false;
}

menuAdminFunction.onclick = function(e) {
  window.location.replace("http://localhost/control-users");
  return false;
}