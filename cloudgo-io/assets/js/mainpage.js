$(document).ready(function() {
    $.ajax({
        url: "/api/mainpage"
    }).then(function(data) {
       $('.username').append(data.username);
    });
});