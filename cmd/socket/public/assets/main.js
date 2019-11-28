$(function() {
    if (!window["WebSocket"]) {
        return;
    }


    var isOK = false
    var content = $("#content");
    var bodys = $("#body");
    var conn = new WebSocket('ws://' + window.location.host + '/ws');

    // Get List Message
    $.ajax({
        url: "http://localhost:4501/message/list",
        method: "GET",
        success: function(data) {
            $.each(data.data, function(key, value) {
                if (value !== null) {
                    bodys.append("<tr><td>"+ value +"</td></tr>")
                }
            })
            isOK = true
        }
    })

    // Input is editable only when socket is opened.
    conn.onopen = function(e) {
        content.attr("disabled", false);
    };

    conn.onclose = function(e) {
        content.attr("disabled", true);
    };

    // Whenever we receive a message, update textarea
    conn.onmessage = function(e) {
        if (e.data != content.val()) {
            if (isOK) {
                bodys.append("<tr><td>"+ e.data +"</td></tr>")
            }
        }
    };

});