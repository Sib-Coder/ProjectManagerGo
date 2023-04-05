$(".delete").click(function (e) {

    let user_id = this.value
    $.ajax({
        type: "POST",
        url: "../api/delete_user/"+user_id,
        statusCode:{
            200:function() {
                $("#"+user_id).remove()
            },
            500: function() {
                let toast = new bootstrap.Toast(document.getElementById('fail'))
                toast.show()
            }
        }
    });
})