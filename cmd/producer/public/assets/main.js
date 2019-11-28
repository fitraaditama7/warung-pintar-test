$(function() {
    if (!window["WebSocket"]) {
        return;
    }
    console.log("coba")

    var content = $("#content");
    var button = $("#send")
    var conn = new WebSocket('ws://localhost:8080/ws');

    // Input is editable only when socket is opened.
    conn.onopen = function(e) {
        content.attr("disabled", false);
    };

    conn.onclose = function(e) {
        content.attr("disabled", true);
    };

    // button trigger when input
    button.on("click", function() {
        conn.send(content.val());
        $.ajax({
            url: 'http://' + window.location.host + '/message/send',
            method: "POST",
            data: {message: content.val()},
            headers: {
                'Content-type': 'application/x-www-form-urlencoded'
            },
            success: function(data) {
                content.val("")
                console.log(data.data)
            },
            error: function(err) {
                console.log(err)
            }
        })
        
    });
});