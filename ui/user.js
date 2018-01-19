function getUserInfo(name) {
    return window.localStorage.getItem(name);
}

function setUserInfo(name, value) {
    window.localStorage.setItem(name, value)
}

function initUser(area, url) {
    let userId = getUserInfo("id")
    let userName = getUserInfo("name")
    if (userId == "" || userName == "") {
        let name = $("#"+ area).val()
        console.log(name)
        return userEnroll(name, url)
    } else {
        console.log(getUserInfo("name"))
        console.log(getUserInfo("id"))
        $("#"+ area).val(getUserInfo("name"))
        return getUserInfo("id")
    }
}

// TODO: 在ui中实现相关接口，避免跨域请求
function userEnroll(name, url) {
    $.ajax({
        type: 'GET',
        url: url + "?name=" + name,
        crossDomain: true,
        jsonp: "callback",
        dataType: "jsonp",
        data: {
            name:name
        },
        success: function (result) {
            console.log("user enroll success")
            notie.alert({type:1, text: "user enroll success", stay: 1.5})
            // $("#"+area).val(result.name)
            setUserInfo("id", result.id)
            setUserInfo("name", result.name)
            return result.id
        },
        error: function (e) {
            notie.alert({type:3, text: e.responseJSON.Message, stay: 1.5})
        }
    })
}