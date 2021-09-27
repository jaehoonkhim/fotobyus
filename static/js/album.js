function createHttp() {
    if (window.ActiveXObject) {
        try {
            return new ActiveXObject("Msxml2.XMLHTTP");
        } catch (e1) {
            try {
                return new ActiveXObject("Microsoft.XMLHTTP");
            } catch (e1) {
                return null;
            }
        }
    } else if (window.XMLHttpRequest) {
        return new XMLHttpRequest();
    }
    return null;
}

function signup() {
    var xmlhttp = createHttp();
    if (!xmlhttp) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var url = "http://" + window.location.host + "/api/signup";
    console.log(url)
    xmlhttp.open("POST", url, true);
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == 4) {
            var json = JSON.parse(xmlhttp.responseText);
            console.log(json.status);
            if (json.status == 0) {
                location.href = "http://" + window.location.host + "/signin";
            } else {
                alert("회원가입 실패")
            }
        }
    }
    xmlhttp.send(signUpData());
}

function signUpData() {
    var request = {};
    request["username"] = encodeURIComponent(document.getElementById("username").value);
    request["userid"] = encodeURIComponent(document.getElementById("userid").value);
    request["email"] = encodeURIComponent(document.getElementById("email").value);
    request["password"] = encodeURIComponent(document.getElementById("pw").value);
    return JSON.stringify(request)
}

function signin() {
    var xmlhttp = createHttp();
    if (!xmlhttp) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var url = "http://" + window.location.host + "/api/signin";
    //console.log(url)
    xmlhttp.open("POST", url, true);
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == 4) {
            var json = JSON.parse(xmlhttp.responseText);
            //console.log(json.status);
            if (json.status == 0) {
                //location.href = "http://" + window.location.host +"/";
                var form = document.createElement("form");
                form.setAttribute("method", "POST");
                form.setAttribute("action", "http://" + window.location.host + "/signin");
                var input = document.createElement("input");
                input.type = "text";
                input.name = 'json';
                input.value = xmlhttp.responseText;
                console.log(xmlhttp.responseText)
                form.appendChild(input);
                document.body.appendChild(form);
                form.submit();
            } else {
                alert("로그인 실패");
            }
        }
    }
    xmlhttp.send(signInData());
}

function signInData() {
    var request = {};
    request["userid"] = encodeURIComponent(document.getElementById("userid").value);
    request["password"] = encodeURIComponent(document.getElementById("pw").value);
    return JSON.stringify(request)
}

function start() {
    var xmlhttp = createHttp();
    if (!xmlhttp) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var url = "http://" + window.location.host + "/api/album";
    xmlhttp.open("POST", url, true);
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == 4) {
            var json = JSON.parse(xmlhttp.responseText);
            if (json.status == 0) {
                var form = document.createElement("form");
                form.setAttribute("method", "POST");
                form.setAttribute("action", "http://" + window.location.host + "/start");
                var input = document.createElement("input");
                input.type = "text";
                input.name = 'json';
                input.value = xmlhttp.responseText;
                console.log(xmlhttp.responseText)
                form.appendChild(input);
                document.body.appendChild(form);
                form.submit();
            } else {
                alert("앨범 생성 실패");
            }
        }
    }
    xmlhttp.send(startData());
}

function startData() {
    var request = {};
    request["title"] = encodeURIComponent(document.getElementById("title").value);
    request["start_date"] = encodeURIComponent(document.getElementById("start_date").value);
    request["end_date"] = encodeURIComponent(document.getElementById("end_date").value);
    request["desc"] = encodeURIComponent(document.getElementById("desc").value);
    request["user_token"] = encodeURIComponent(document.getElementById("token").value);
    return JSON.stringify(request)
}

function profile(token) {
    var xmlhttp = createHttp();
    if (!xmlhttp) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var url = "http://" + window.location.host + "/api/profile?user=" + token;
    xmlhttp.open("GET", url, true);
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == 4) {
            var json = JSON.parse(xmlhttp.responseText);
            if (json.status == 0) {
                document.getElementById("name").innerHTML = decodeURIComponent(json.name);
                document.getElementById("email").innerHTML = decodeURIComponent(json.email);
                var txt = "";
                txt = "<ul>\n";
                for (var i = 0; i < json.albums.length; ++i) {
                    var terms = decodeURIComponent(json.albums[i].start_date) + "~" + decodeURIComponent(json.albums[i].end_date);
                    if (terms == "~") {
                        terms = "";
                    }
                    txt += "<li><div><a href=\"/album/" + decodeURIComponent(json.albums[i].albumid) + "\"><b>" + decodeURIComponent(json.albums[i].title) + "</b></a><br/>" + terms + "<br/>" + decodeURIComponent(json.albums[i].desc) + "<br/><br/></div></li>\n";
                }

                txt += "</ul>"
                document.getElementById("albums").innerHTML = txt;

            } else {
                alert("프로필 조회 실패");
            }
        }
    }
    xmlhttp.send();

}


function profileData(token) {
    var request = {};
    request["token"] = token;
    return JSON.stringify(request)
}

function album(albumid, user) {
    var xmlhttp = createHttp();
    if (!xmlhttp) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var url = "http://" + window.location.host + "/api/album?id=" + albumid + "&user=" + user;
    xmlhttp.open("GET", url, true);
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == 4) {
            var json = JSON.parse(xmlhttp.responseText);
            if (json.status == 0) {
                document.getElementById("title").innerHTML = "<h3>" + decodeURIComponent(json.title) + "</h3>";
                terms = decodeURIComponent(json.start_date) + "~" + decodeURIComponent(json.end_date);
                if (terms == "~") {
                    document.getElementById("date").innerHTML = "";
                } else {
                    document.getElementById("date").innerHTML = terms;
                }
                document.getElementById("desc").innerHTML = decodeURIComponent(json.desc);
            } else {
                alert("앨범 조회 실패");
            }
        }
    }
    xmlhttp.send();
}

function albumSettings(albumid, user) {
    var xmlhttp = createHttp();
    if (!xmlhttp) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var url = "http://" + window.location.host + "/api/album?id=" + albumid + "&user=" + user;
    xmlhttp.open("GET", url, true);
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == 4) {
            var json = JSON.parse(xmlhttp.responseText);
            if (json.status == 0) {
                document.getElementById("title").value = decodeURIComponent(json.title);
                document.getElementById("start_date").value = decodeURIComponent(json.start_date);
                document.getElementById("end_date").value = decodeURIComponent(json.end_date);
                document.getElementById("desc").value = decodeURIComponent(json.desc);
            } else {
                alert("앨범 조회 실패");
            }
        }
    }
    xmlhttp.send();
}

function albumUpdate() {
    console.log("albumUpdate");
    var xhr = createHttp();
    if (!xhr) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var url = "http://" + window.location.host + "/api/album";
    xhr.open("PUT", url, true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            var json = JSON.parse(xhr.responseText);
            if (json.status == 0) {
                location.href = "/album/" + json.albumid;
            } else {
                alert("blah");
            }
        }
    }

    xhr.send(albumUpdateData());
}

function albumUpdateData() {
    var request = {};
    request["title"] = encodeURIComponent(document.getElementById("title").value);
    request["start_date"] = encodeURIComponent(document.getElementById("start_date").value);
    request["end_date"] = encodeURIComponent(document.getElementById("end_date").value);
    request["desc"] = encodeURIComponent(document.getElementById("desc").value);
    request["user_token"] = encodeURIComponent(document.getElementById("token").value);
    request["album_id"] = encodeURIComponent(document.getElementById("albumid").value);
    return JSON.stringify(request)

}

function upload() {
    var xhr = createHttp();
    if (!xhr) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var uploadBuf = document.getElementById("uploadForm"); //폼 자체를 전송
    var formData = new FormData(uploadBuf)

    var url = "http://" + window.location.host + "/api/uploadimage";
    xhr.open("POST", url, true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            console.log(xhr.responseText);
            var json = JSON.parse(xhr.responseText);
            if (json.status == 0) {
                //console.log("업로드 성공1");
                document.getElementById("image").value = "";
                images(document.getElementById("albumid").value, document.getElementById("userkey").value);
            } else {
                alert("업로드 실패");
            }
        }
    }

    xhr.send(formData);
}

function images(albumID, user) {
    console.log("images")
    var xhr = createHttp();
    var url = "http://" + window.location.host + "/api/image?id=" + albumID + "&user=" + user;
    console.log(url)
    xhr.open("GET", url, true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            var json = JSON.parse(xhr.responseText);
            containerTag = document.getElementById("imageContainer");
            while (containerTag.hasChildNodes()) {
                containerTag.removeChild(containerTag.firstChild)
            }
            console.log(json);
            if (json.keys == null) {
                return;
            }

            for (var i = 0; i < json.keys.length; i++) {
                var divTag = document.createElement("div");
                divTag.setAttribute("class", "thumb");
                divTag.setAttribute("style", "float:left;margin:0px 5px 5px 0px;border:1px solid #999;");
                var aTag = document.createElement("a");
                aTag.href = "/detail/" + albumID + "/" + json.keys[i];
                var imgTag = document.createElement("img");
                imgTag.style.width = "auto";
                imgTag.style.height = "200px";
                imgTag.src = "/api/imageview?blobKey=" + json.keys[i];
                var delButton = document.createElement("input");
                delButton.setAttribute("value", json.keys[i]);
                delButton.setAttribute("style", "float:left;width:10px;height:10px;");
                delButton.onclick = function() {
                    delImage(this.value, albumID, user);
                }
                delButton.src = "/img/11968905106_b20222983a_q.jpg";
                delButton.type = "image";

                var editButton = document.createElement("input");
                editButton.setAttribute("style", "float:right;width:10px;height:10px;");
                editButton.onclick = function() {
                    editImage(this.value, albumID, user);
                }
                editButton.src = "/img/edit-1103598_640.png";
                editButton.type = "image";

                var div2Tag = document.createElement("div");
                div2Tag.setAttribute("style", "position:relative;height:10px;width:auto;top:0px;left:0px;");

                aTag.appendChild(imgTag);

                divTag.appendChild(div2Tag);
                divTag.appendChild(aTag);
                div2Tag.appendChild(delButton);
                div2Tag.appendChild(editButton);
                containerTag.appendChild(divTag);
            }

        }
    }
    xhr.send();

}

function editImage(blobKe, id, user) {
    var xhr = createHttp();
    if (!xhr) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }
    var url = "http://" + window.location.host + "/api/editor?blobkey=" + blobKey + "&id=" + id + "&user=" + user;
    xhr.open("GET", url, true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            console.log(xhr.responseText);
            var json = JSON.parse(xhr.responseText);
            if (json.status == 0) {
                //images(json.albumid, user);
            } else {
                alert("편집기 시작 실패");
            }
        }
    }

    xhr.send();

}

function delImage(blobKey, id, user) {
    var xhr = createHttp();
    if (!xhr) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var url = "http://" + window.location.host + "/api/image?blobkey=" + blobKey + "&id=" + id + "&user=" + user;
    xhr.open("DELETE", url, true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            console.log(xhr.responseText);
            var json = JSON.parse(xhr.responseText);
            if (json.status == 0) {
                images(json.albumid, user);
            } else {
                alert("삭제 실패");
            }
        }
    }

    xhr.send();
}

function gotoAlbumSettings(albumid, user) {
    console.log("gotoAlbumSettings");
    location.href = "/album/" + albumid + "/settings";
    return false;
}

function detail(imageID, user) {
    var content = document.getElementById("content");
    var imgTag = document.createElement("img");
    imgTag.setAttribute("style", "max-width:450px;")
    imgTag.src = "/api/imageview?blobKey=" + imageID;
    content.appendChild(imgTag);

    //comment
    //var comment = document.getElementById("list_comment");
    getComment(imageID);
}

function getComment(imageID) {
    var xhr = createHttp();
    if (!xhr) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var url = "http://" + window.location.host + "/api/comment?id=" + imageID;
    xhr.open("GET", url, true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            console.log(xhr.responseText);
            var json = JSON.parse(xhr.responseText);
            if (json.status == 0) {
                console.log(json)
                var comment = document.getElementById("list_comment");
                while (comment.hasChildNodes()) {
                    comment.removeChild(comment.firstChild)
                }
                for (var i = 0; i < json.comments.length; i++) {
                    cmnt = json.comments[i];
                    var container = document.createElement("div");
                    container.id = "item_comment"
                    container.setAttribute("style", "overflow:hidden")
                    container.innerHTML = "<hr><p>" + decodeURIComponent(cmnt.commentvalue) + "</p>";
                    comment.appendChild(container);
                }

            } else {
                alert("댓글 검색 실패");
            }
        }
    }

    xhr.send(commentData());
}

function sendComment() {
    var xhr = createHttp();
    if (!xhr) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var url = "http://" + window.location.host + "/api/comment";
    xhr.open("POST", url, true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            console.log(xhr.responseText);
            var json = JSON.parse(xhr.responseText);
            if (json.status == 0) {
                document.getElementById("comment_value").value = "";
                getComment(document.getElementById("imageid").value);
            } else {
                alert("업로드 실패");
            }
        }
    }

    xhr.send(commentData());
}

function commentData() {
    var request = {};
    var cmnt = document.getElementById("comment_value").value;
    cmnt = cmnt.replace(/\r?\n/g, '<br/>');
    request["commentvalue"] = encodeURIComponent(cmnt);
    request["usertoken"] = encodeURIComponent(document.getElementById("usertoken").value);
    request["albumid"] = encodeURIComponent(document.getElementById("albumid").value);
    request["imageid"] = encodeURIComponent(document.getElementById("imageid").value);
    console.log(request);
    return JSON.stringify(request)
}

function gotoProfileSettings(user) {
    location.href = "/profile/settings";
    return false;
}

function gotoProfile() {
    location.href = "/profile";
    return false;
}

function getProfileSettings(user) {
    var xmlhttp = createHttp();
    if (!xmlhttp) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var url = "http://" + window.location.host + "/api/profile?user=" + user;
    xmlhttp.open("GET", url, true);
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == 4) {
            var json = JSON.parse(xmlhttp.responseText);
            if (json.status == 0) {
                document.getElementById("username").value = decodeURIComponent(json.name);
                document.getElementById("email").value = decodeURIComponent(json.email);
                document.getElementById("userid").value = decodeURIComponent(json.id);
                document.getElementById("userid").disabled = true;
            } else {
                alert("프로필 조회 실패");
            }
        }
    }
    xmlhttp.send();

}

function updateProfileSettings() {
    console.log("updateProfileSettings");
    var xmlhttp = createHttp();
    if (!xmlhttp) {
        alert('Cannot create a XMLHTTP instance');
        return;
    }

    var url = "http://" + window.location.host + "/api/profile";
    xmlhttp.open("PUT", url, true);
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == 4) {
            var json = JSON.parse(xmlhttp.responseText);
            if (json.status == 0) {
                location.href = "http://" + window.location.host + "/profile";
            } else {
                alert("프로필 설정 실패");
            }
        }
    }
    xmlhttp.send(profileSettingsData())
}

function profileSettingsData() {
    var request = {};
    request["username"] = encodeURIComponent(document.getElementById("username").value);
    request["email"] = encodeURIComponent(document.getElementById("email").value);
    request["usertoken"] = encodeURIComponent(document.getElementById("usertoken").value);
    return JSON.stringify(request)
}
