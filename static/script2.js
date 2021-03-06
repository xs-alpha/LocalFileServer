const form = document.querySelector("form"),
    fileInput = document.querySelector(".file-input"),
    progressArea = document.querySelector(".progress-area"),
    uploadedArea = document.querySelector(".uploaded-area");


// form click event
form.addEventListener("click", () => {
    fileInput.click();
});

fileInput.onchange = ({ target }) => {
    var files = target.files;
    for (var i = 0, f; f = files[i]; i++) {
        var fn = new FormData;
        fn.append("filename", f);
        console.log(f);
        // file upload function
        uploadFile(fn);
        function uploadFile(name) {
            let urluploadlocal = document.location.href + "upload"; //getting url of upload local
            console.log(urluploadlocal)
            let xhr = new XMLHttpRequest(); //creating new xhr object (AJAX)
            xhr.open("post", urluploadlocal, true); //sending post request to the specified URL
            xhr.upload.addEventListener("progress", ({ loaded, total }) => { //file uploading progress event
                let fileLoaded = Math.floor((loaded / total) * 100);  //getting percentage of loaded file size
                let fileTotal = Math.floor(total / 1000); //gettting total file size in KB from bytes
                let fileSize;
                // if file size is less than 1024 then add only KB else convert this KB into MB
                (fileTotal < 1024) ? fileSize = fileTotal + " KB" : fileSize = (loaded / (1024 * 1024)).toFixed(2) + " MB";
                let progressHTML = `<li class="row">
                            <i class="fas fa-file-alt"></i>
                            <div class="content">
                              <div class="details">
                                <span class="name">${name} • Uploading</span>
                                <span class="percent">${fileLoaded}%</span>
                              </div>
                              <div class="progress-bar">
                                <div class="progress" style="width: ${fileLoaded}%"></div>
                              </div>
                            </div>
                          </li>`;
                // uploadedArea.innerHTML = ""; //uncomment this line if you don't want to show upload history
                uploadedArea.classList.add("onprogress");
                progressArea.innerHTML = progressHTML;
                if (loaded == total) {
                    progressArea.innerHTML = "";
                    let uploadedHTML = `<li class="row">
                              <div class="content upload">
                                <i class="fas fa-file-alt"></i>
                                <div class="details">
                                  <span class="name">${name} • Uploaded</span>
                                  <span class="size">${fileSize}</span>
                                </div>
                              </div>
                              <i class="fas fa-check"></i>
                            </li>`;
                    uploadedArea.classList.remove("onprogress");
                    // uploadedArea.innerHTML = uploadedHTML; //uncomment this line if you don't want to show upload history
                    uploadedArea.insertAdjacentHTML("afterbegin", uploadedHTML); //remove this line if you don't want to show upload history
                }
            });
            xhr.onload = function (e) {
                if (this.status == 200 || this.status == 304) {
                    layer.msg("上传成功");
                    // alert("上传成功");
                    console.log(this.responseText);
                } else {
                    layer.msg("上传失败");
                    // alert("上传失败");
                    console.log(this.responseText);
                }
            }
            // 请求结束
            xhr.onloadend = e => {
                layer.msg("上传结束");
                console.log('request loadend');
            };
            // 请求超时
            xhr.ontimeout = e => {
                layer.msg("请求超时");
                console.log('request timeout');
            };
            xhr.onerror = function (e) {
                // layer.msg("上传失败");
                console.log(this.responseText);
            }

            xhr.send(name); //sending form data
        }

    }
}

