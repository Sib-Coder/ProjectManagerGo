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