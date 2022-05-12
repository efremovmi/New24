const timeOutMessage = 2000
const message = document.querySelector(".message")
const editFunction = document.getElementById("edit_button")
const exitFunction = document.getElementById("return_to_main_menu")

exitFunction.onclick = function(e) {
  window.location.replace("http://localhost/news");
  return false
}



editFunction.onclick = function(e) {


  const formData = new FormData();
  formData.append('id', document.getElementById('id-news').value);
  formData.append('header', document.getElementById('name-news').value);
  formData.append('news', document.getElementById('full-news').value);
  formData.append('image', document.getElementById('custom-image-upload').files[0]);


  if (document.getElementById('create').checked == true) {
    axios.post(`http://localhost/news/save-post`, formData, {timeout: 5000})
    .then(function (response) {
      if (response.statusText == "OK") {
        message.innerHTML = "Новость добавлена!"
        message.classList.add("active")
        setTimeout(() => {
          message.classList.remove("active")
        }, timeOutMessage)

        document.getElementById('name-news').value = "";
        document.getElementById('full-news').value = ""; 

         axios.get("http://localhost/news/get-all-news", {timeout: 5000})
        .then(function (response) {
          if (response.statusText == "OK") {
            div = document.getElementById('tbody');
            div.innerHTML = ``;
            for (let i = 0; i < response.data.news_list.length; i += 1) {
              console.log(response.data.news_list[i]);

            div.innerHTML += `<tr>
                              <td class="text-left">`+response.data.news_list[i].id+`</td>
                              <td class="text-left"><a href="http://localhost/news/get-post?header=`+response.data.news_list[i].header+`">`+response.data.news_list[i].header+`</a></td>                               
                              <td class="text-left">`+response.data.news_list[i].news+`</td>
                              <td class="text-left"><img src="http:/`+`/localhost/` + response.data.news_list[i].header + `/` + response.data.news_list[i].header + `.jpeg" alt=""></td>
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
          console.log(error);
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
      console.log(error);
      message.innerHTML = error;
      message.classList.add("error", "active")
      setTimeout(() => {
       message.classList.remove("error", "active")
      }, timeOutMessage)
    });

  }else if (document.getElementById('delete').checked == true) {
    axios.post(`http://localhost/news/delete-post`, formData, {timeout: 5000})
    .then(function (response) {
      if (response.statusText == "OK") {
        message.innerHTML = "Новость удалена!"
        message.classList.add("active")
        setTimeout(() => {
          message.classList.remove("active")
        }, timeOutMessage)


        document.getElementById('name-news').value = "";
        document.getElementById('full-news').value = ""; 

         axios.get("http://localhost/news/get-all-news", {timeout: 5000})
        .then(function (response) {
          if (response.statusText == "OK") {
            div = document.getElementById('tbody');
            div.innerHTML = ``;
            for (let i = 0; i < response.data.news_list.length; i += 1) {
              console.log(response.data.news_list[i]);

              div.innerHTML += `<tr>
                                <td class="text-left">`+response.data.news_list[i].id+`</td>
                                <td class="text-left"><a href="http://localhost/news/get-post?header=`+response.data.news_list[i].header+`">`+response.data.news_list[i].header+`</a></td>                               
                                <td class="text-left">`+response.data.news_list[i].news+`</td>
                                <td class="text-left"><img src="http:/`+`/localhost/` + response.data.news_list[i].header + `/` + response.data.news_list[i].header + `.jpeg" alt=""></td>
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
          console.log(error);
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
      console.log(error);
      message.innerHTML = error;
      message.classList.add("error", "active")
      setTimeout(() => {
       message.classList.remove("error", "active")
      }, timeOutMessage)
    });
  }
 
  else if (document.getElementById('update').checked == true) {
    axios.post(`http://localhost/news/update-post`, formData, {timeout: 5000})
    .then(function (response) {
      if (response.statusText == "OK") {
        message.innerHTML = "Новость обновлена!"
        message.classList.add("active")
        setTimeout(() => {
          message.classList.remove("active")
        }, timeOutMessage)

        document.getElementById('name-news').value = "";
        document.getElementById('full-news').value = ""; 

         axios.get("http://localhost/news/get-all-news", {timeout: 5000})
        .then(function (response) {
          if (response.statusText == "OK") {
            div = document.getElementById('tbody');
            div.innerHTML = ``;
            for (let i = 0; i < response.data.news_list.length; i += 1) {
              console.log(response.data.news_list[i]);

              div.innerHTML += `<tr>
                                <td class="text-left">`+response.data.news_list[i].id+`</td>
                                <td class="text-left"><a href="http://localhost/news/get-post?header=`+response.data.news_list[i].header+`">`+response.data.news_list[i].header+`</a></td>                               
                                <td class="text-left">`+response.data.news_list[i].news+`</td>
                                <td class="text-left"><img src="http:/`+`/localhost/` + response.data.news_list[i].header + `/` + response.data.news_list[i].header + `.jpeg" alt=""></td>
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
          console.log(error);
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
      console.log(error);
      message.innerHTML = error;
      message.classList.add("error", "active")
      setTimeout(() => {
       message.classList.remove("error", "active")
      }, timeOutMessage)
    });
  }
  return false
}
