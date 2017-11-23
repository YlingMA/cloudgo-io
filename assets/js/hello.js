$(document).ready(function() {
    $.ajax({
        url: "/api/test"
    }).then(function(data) {
       $('.name').append(data.name);
       $('.q1').append(data.q1);
       $('.q2').append(data.q2);
       $('.q3').append(data.q3);
       $('.q3last').append(data.q3last);
    });
});
