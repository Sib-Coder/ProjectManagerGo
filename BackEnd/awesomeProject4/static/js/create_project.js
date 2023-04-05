$("#update").click(function (e) {

    $('#title').removeClass("is-invalid");


    $('#title').addClass("is-valid");

    let title_name = $("#title").val();
    const re = new  RegExp('(^\\s+$|^$)')
    is_valid = true;
    if (re.test(title_name)) {
        is_valid = false
        $('#title').removeClass("is-valid");
        $("#title").addClass("is-invalid")
        // show error
    }

    let description_val = $("#description").val();
    let color_val = $("#color").val();
    let text_color_val = $("#text_color").val();
    console.log(color_val)
    let send_data = JSON.stringify({
        Name: title_name,
        description: description_val,
        Colour: color_val,
        TextColor: text_color_val
    })

    if(!is_valid){
        return
    }

    $.ajax({
        type: "POST",
        url: "../../../api/new_project/",
        data: send_data,
        contentType: "application/json",
        dataType: "json",
        statusCode:{
            200:function() {
                let toast = new bootstrap.Toast(document.getElementById('done'))
                toast.show()
            },
            500: function() {
                let toast = new bootstrap.Toast(document.getElementById('fail'))
                toast.show()
            }
        }
    });
    
});