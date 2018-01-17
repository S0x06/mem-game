function getUserId() {
    let id = window.localStorage.getItem("id");
    if (id == null) {
        return createUser()
    }
    return id
}

function createUser() {
    let id = guid()
    window.localStorage.setItem("id", id);
    return id
}

function guid() {
    function s4() {
        return Math.floor((1 + Math.random()) * 0x10000)
            .toString(16)
            .substring(1);
    }
    return s4() + s4() + '-' + s4() + '-' + s4() + '-' +
        s4() + '-' + s4() + s4() + s4();
}

function initUser(area, url) {
    let id = getUserId()
    $("#"+area).val(id)
    userEnroll(id, url)
}

// TODO: 在ui中实现相关接口，避免跨域请求
function userEnroll(id, url) {
    $.ajax({
        type: 'GET',
        url: url,
        crossDomain: true,
        jsonp: "callback",
        dataType: "jsonp",
        data: {
            id:id
        },
        success: function (result) {
            console.log("user enroll success")
        },
        error: function (e) {
            notie.alert({type:1, text: e.responseJSON.Message, stay: 1.5})
        }
    })
}