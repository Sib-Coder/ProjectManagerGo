document.addEventListener('DOMContentLoaded', function() {
    var calendarEl = document.getElementById('calendar');
    
    addr = window.location.href.split('/')
    id = addr[addr.length - 1]
    arr = []
    $.ajax({
        type: "GET",
        url: "../api/tasks/"+id,
        success: function (response) {
            if (response != null) {
                for (let i = 0; i < response.length; i++) {
                    arr[i] = response[i];
                }
            }
        }
    })
    .always(function() {
        if (arr.length > 0 ) {
            var calendar = new FullCalendar.Calendar(calendarEl, {
                initialView: 'dayGridMonth',
                themeSystem: 'standard',
                events: arr
            });
        } else {
            var calendar = new FullCalendar.Calendar(calendarEl, {
                initialView: 'dayGridMonth'
            });
        }
        calendar.render();
    });
});






$("#delete").click(function (e) {

    let page_url = window.location.href.split('/')
    let project_id = page_url[page_url.length - 1]
    $.ajax({
        type: "POST",
        url: "../../api/delete_project/"+project_id,
        statusCode:{
            200:function() {
                let toast = new bootstrap.Toast(document.getElementById('done_del'))
                toast.show()
                location.reload()
            },
            500: function() {
                let toast = new bootstrap.Toast(document.getElementById('fail'))
                toast.show()
            }
        }
    });
});