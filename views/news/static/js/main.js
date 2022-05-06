const timeOutMessage = 2000;
const exitFunction = document.getElementById("exit");
const menuAdminFunction = document.getElementById("menu_admin");



exitFunction.onclick = function(e) {
  document.cookie = "token= ;path=/";
  window.location.replace("http://localhost:8000/auth");
  return false
}

menuAdminFunction.onclick = function(e) {
  window.location.replace("http://localhost:8001/control-users");
  return false
}
